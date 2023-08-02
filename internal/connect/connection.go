package connect

import (
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

type Job struct {
	IP        string `env:"SSHME_IP"`
	User      string `env:"SSHME_USER"`
	PathToKey string `env:"SSHME_KEY"`
	Command   string `env:"-"`
	Answer    []byte `env:"-"`
}

// StartJob return Job structure
func StartJob() (*Job, error) {
	return &Job{Answer: make([]byte, 0)}, nil
}

// ExecRemotely establish connection and run command
func (j *Job) ExecRemotely() ([]byte, error) {

	privateBytes, err := os.ReadFile(j.PathToKey)
	if err != nil {
		log.Println("Failed to load private key: ", err)
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Println("Failed to parse private key: ", err)
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: j.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", j.IP, config)
	if err != nil {
		log.Println("Dial failed:", err)
		return nil, err
	}

	defer client.Close()

	session, err := client.NewSession()

	if err != nil {
		log.Println("Session failed:", err)
		return nil, err
	}

	output, err := session.CombinedOutput(j.Command)

	if err != nil {
		log.Println("CombinedOutput failed:", err)
		return nil, err
	}

	return output, nil
}
