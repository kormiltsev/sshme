package connect

import (
	"flag"
	"fmt"
	"log"
	"os"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// ParseFlags parses flags and environments.
func (j *Job) ParseFlags() error {

	j.file()

	if err := j.environment(); err != nil {
		log.Println("Environment error:", err)
	}

	flag.StringVar(&j.IP, "ip", j.IP, fmt.Sprintf("server's IP (%s)", j.IP))
	flag.StringVar(&j.User, "u", j.User, fmt.Sprintf("username (%s)", j.User))
	flag.StringVar(&j.PathToKey, "k", j.PathToKey, fmt.Sprintf("path to key (ex: '/Users/username/.ssh/id_rsa' ) (%s)", j.PathToKey))
	flag.StringVar(&j.Command, "exec", "", "remote server to exec")

	flag.Parse()

	if j.IP == "" || j.User == "" || j.PathToKey == "" {
		return fmt.Errorf("no arguments")
	}
	return nil
}

// Environment returns ENV values
func (j *Job) environment() error {
	err := env.Parse(j)
	if err != nil {
		return err
	}
	// log.Println("got from env:", j)
	return nil
}

// File returns values from .env file
func (j *Job) file() {
	godotenv.Load()
	j.IP = os.Getenv("SSHME_IP")
	j.User = os.Getenv("SSHME_USER")
	j.PathToKey = os.Getenv("SSHME_KEY")
}
