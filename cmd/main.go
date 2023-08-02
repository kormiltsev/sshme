package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kormiltsev/sshme/internal/connect"
)

func main() {

	// get Job structure
	job, err := connect.StartJob()
	if err != nil {
		log.Println("Creation error: ", err)
		os.Exit(1)
	}

	// get settings from flags, ENV and .env-file
	err = job.ParseFlags()
	if err != nil {
		log.Println("ParseFlags error: ", err)
		os.Exit(1)
	}

	// put required command here
	if job.Command == "" {
		command := `echo "===================== RAM ====================="; free -h | awk 'NR==1 {print $1 "  " $2 "  " $3 " " $4 " " $5}'; free -h | awk 'NR==2 {print $2 " " $3 "  " $4 "  " $5 " " $6}'; echo "==================== DRIVES ==================="; df -H`

		err = job.SetCommand(command)
		if err != nil {
			log.Println("SetCommand error: ", err)
			os.Exit(1)
		}
	}

	// run command on remote server
	answer, err := job.ExecRemotely()
	if err != nil {
		log.Println("ExecRemotely error: ", err)
		os.Exit(1)
	}

	// do some other operations with answer
	//
	fmt.Println(string(answer))
	//
	// ===================================
}
