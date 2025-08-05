package flags

import (
	"flag"
	"fmt"
	"goble/api"
	"os"
)

func ParseFlags() api.Flags {
	user := flag.String("username", "", "The username that is used for the ssh connection")
	password := flag.String("password", "", "Password for the ssh login")
	hostsfile := flag.String("hostfile", "", "Path to the hosts file")

	flag.Parse()

	if *user == "" || *password == "" {
		fmt.Println("Username and password flags are required.")
		flag.Usage()
		os.Exit(1)
	}
	return api.Flags{
		User:      *user,
		Password:  *password,
		Hostsfile: *hostsfile,
	}
}
