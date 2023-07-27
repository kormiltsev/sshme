package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kormiltsev/sshme/internal/connect"
)

func main() {

	job, err := connect.StartJob()
	if err != nil {
		log.Println("Creation error: ", err)
		os.Exit(1)
	}

	err = job.ParseFlags()
	if err != nil {
		log.Println("ParseFlags error: ", err)
		os.Exit(1)
	}

	command := `echo "==================== RAM ===================="; free -h | awk 'NR==1 {print $1 "  " $2 "  " $3}'; free -h | awk 'NR==2 {print $2 " " $3 " " $4}'; echo "=================== DRIVES =================="; df -H`

	err = job.SetCommand(command)
	if err != nil {
		log.Println("SetCommand error: ", err)
		os.Exit(1)
	}

	answer, err := job.ExecRemotely()
	if err != nil {
		log.Println("ExecRemotely error: ", err)
		os.Exit(1)
	}

	fmt.Println(answer)

	// do some other operations with answer
	//
	// ===================================
}
