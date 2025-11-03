package main

import (
	commandsParser "goble/parser/commands"
	flagParser "goble/parser/flags"
	hostParser "goble/parser/hosts"
	"goble/runner"
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)
	flags := flagParser.ParseFlags()

	hosts, err := hostParser.ParseHosts(flags.HostsFile)
	if err != nil {
		log.Fatalf("Error during hosts parsing")
	}
	commands, err := commandsParser.ParseCommands(flags.Commands)
	if err != nil {
		log.Fatalf("Error during commands parsing")
	}

	var wg sync.WaitGroup
	runner.ProcessHosts(hosts, commands, flags.User, flags.Password, &wg)

	log.Println("\nAll connections have been processed.")
}
