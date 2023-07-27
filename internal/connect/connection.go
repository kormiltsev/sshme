package connect

import (
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

type Job struct {
	IP        string
	User      string
	PathToKey string
	Command   string
	Answer    []byte
}

func StartJob() (*Job, error) {
	return &Job{Answer: make([]byte, 0)}, nil
}

func ExecRemotely(ip, user, pathToKey string) ([]byte, error) {

	privateBytes, err := os.ReadFile(pathToKey)
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
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", ip, config)
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

	command := job()
	output, err := session.CombinedOutput(command)

	if err != nil {
		log.Println("CombinedOutput failed:", err)
		return nil, err
	}

	return output, nil
}
