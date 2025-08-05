package hosts

import (
	"goble/api"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func ParseHosts(hostfile string) (api.Hosts, error) {
	yamlFile, err := os.ReadFile(hostfile)
	if err != nil {
		log.Fatalf("Failed to read hosts.yaml: %v", err)
	}
	var hosts api.Hosts
	err = yaml.Unmarshal(yamlFile, &hosts)
	if err != nil {
		return api.Hosts{}, err
	}
	return hosts, nil
}
