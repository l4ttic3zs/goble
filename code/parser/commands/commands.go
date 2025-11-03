package commands

import (
	"goble/api"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func ParseCommands(commandsFile string) (api.Commands, error) {
	yamlFile, err := os.ReadFile(commandsFile)
	if err != nil {
		log.Fatalf("Failed to read commands.yaml: %v", err)
	}
	var commands api.Commands
	err = yaml.Unmarshal(yamlFile, &commands)
	if err != nil {
		return api.Commands{}, err
	}
	return commands, nil
}
