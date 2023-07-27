package connect

import (
	"flag"
	"fmt"
	"log"
	"os"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// parseflag parses flags.
func (j *Job) ParseFlags() error {

	if err := j.Environment(); err != nil {
		return err
	}

	flag.StringVar(&j.IP, "ip", "", "server's IP")
	flag.StringVar(&j.User, "u", "", "username")
	flag.StringVar(&j.PathToKey, "k", "", "path to key (ex: '/Users/username/.ssh/id_rsa' )")
	flag.StringVar(&j.PathToKey, "exec", "", "remote server to exec")

	flag.Parse()

	if j.IP == "" || j.User == "" || j.PathToKey == "" {
		return fmt.Errorf("no arguments")
	}
	return nil
}

// Environment returns ENV values
func (j *Job) Environment() error {
	err := env.Parse(j)
	if err != nil {
		return err
	}
	log.Println("got from env:", j)
	return nil
}

func (j *Job) File() {
	godotenv.Load()
	j.IP = os.Getenv("SSHME_IP")
	j.User = os.Getenv("SSHME_USER")
	j.PathToKey = os.Getenv("SSHME_KEY")
}
