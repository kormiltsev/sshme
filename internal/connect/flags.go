package connect

import (
	"flag"
	"fmt"
)

// parseflag parses flags.
func (j *Job) ParseFlags() error {

	// Client settings
	flag.StringVar(&j.IP, "ip", "", "server's IP")

	flag.StringVar(&j.User, "u", "", "username")

	flag.StringVar(&j.PathToKey, "k", "", "path to key (ex: '/Users/username/.ssh/id_rsa' )")

	flag.Parse()

	if j.IP == "" || j.User == "" || j.PathToKey == "" {
		return fmt.Errorf("no arguments")
	}
	return nil
}
