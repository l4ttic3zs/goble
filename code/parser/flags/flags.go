package flags

import (
	"flag"
	"goble/api"
	"log"
	"os"
)

func ParseFlags() api.Flags {
	user := flag.String("username", "", "The username that is used for the ssh connection")
	password := flag.String("password", "", "Password for the ssh login")
	hostsFile := flag.String("hostfile", "", "Path to the hosts file")
	commandsFile := flag.String("commands", "", "Path to the commands file")

	flag.Parse()

	if *user == "" || *password == "" {
		log.Println("Username and password flags are required.")
		flag.Usage()
		os.Exit(1)
	} else if *hostsFile == "" {
		log.Println("No host file provided, nothing to do...")
		flag.Usage()
		os.Exit(1)
	} else if *commandsFile == "" {
		log.Println("No commands given, nothing to run...")
		flag.Usage()
		os.Exit(1)
	}
	return api.Flags{
		User:      *user,
		Password:  *password,
		HostsFile: *hostsFile,
		Commands:  *commandsFile,
	}
}
