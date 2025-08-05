package main

import (
	"fmt"
	flagParser "goble/parser/flags"
	hostParser "goble/parser/hosts"
	"goble/runner"
	"log"
	"sync"
)

func main() {

	flags := flagParser.ParseFlags()

	hosts, err := hostParser.ParseHosts(flags.Hostsfile)
	if err != nil {
		log.Fatalf("Error during hosts parsing")
	}

	var wg sync.WaitGroup
	for _, host := range hosts.HostsConfigs {
		wg.Add(1)
		go runner.RunSSH(host, flags.User, flags.Password, &wg)
	}

	wg.Wait()
	fmt.Println("\nAll connections have been processed.")
}
