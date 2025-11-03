package runner

import (
	"fmt"
	"goble/api"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func ProcessHosts(hosts api.Hosts, commands api.Commands, user, password string, wg *sync.WaitGroup) {
	processGroup(hosts.Groups, commands.Command, user, password, wg)

	log.Println("All SSH operations completed.")
}

func processGroup(groups []api.Group, commands []string, user, password string, wg *sync.WaitGroup) {
	log.Printf("asd")
	for _, group := range groups {
		log.Printf("asdasd")
		if len(group.SubGroups) > 0 {
			processGroup(group.SubGroups, commands, user, password, wg)
		} else {
			for _, host := range group.Hosts {
				wg.Add(1)
				go runSSH(host, commands, user, password, wg)
			}
		}
	}
	wg.Wait()
}

func runSSH(hostConfig api.HostConfig, commands []string, user, password string, wg *sync.WaitGroup) {
	defer wg.Done()
	outputDir := hostConfig.Name

	address := fmt.Sprintf("%s:%d", hostConfig.IP, hostConfig.Port)
	log.Printf("Connecting to %s...", address)

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", address, config)
	if err != nil {
		log.Printf("Error connecting to %s: %v", address, err)
		return
	}
	defer conn.Close()

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Printf("Error creating directory %s for host %s: %v", outputDir, hostConfig.Name, err)
	}

	for _, cmd := range commands {
		session, err := conn.NewSession()
		if err != nil {
			log.Printf("Error creating session for command '%s' on %s: %v", cmd, address, err)
			continue
		}
		defer session.Close()

		safeCmd := strings.ReplaceAll(cmd, " ", "_")
		fileName := filepath.Join(outputDir, fmt.Sprintf("%s_output.txt", safeCmd))

		output, err := session.CombinedOutput(cmd)
		if err != nil {
			log.Printf("Error running command '%s' on %s: %v", cmd, address, err)
			continue
		}
		fileContent := fmt.Sprintf("%s\n", string(output))

		if writeErr := os.WriteFile(fileName, []byte(fileContent), 0644); writeErr != nil {
			log.Printf("Error writing output to file %s for host %s: %v", fileName, hostConfig.Name, writeErr)
		} else {
			log.Printf("Command '%s' output saved to %s", cmd, fileName)
		}
		log.Printf("\n------------------ Output from %s for command '%s' ------------------\n%s", address, cmd, string(output))
	}

	log.Printf("------------------ All commands finished on %s ------------------\n", address)
}
