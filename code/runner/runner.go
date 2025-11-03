package runner

import (
	"fmt"
	"goble/api"
	"log"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func ProcessHosts(hosts api.Hosts, user, password string, wg *sync.WaitGroup) {
	processGroup(hosts.Groups, user, password, wg)

	wg.Wait()
	log.Println("All SSH operations completed.")
}

func processGroup(groups []api.Group, user, password string, wg *sync.WaitGroup) {
	for _, group := range groups {
		for _, host := range group.Hosts {
			wg.Add(1)
			go runSSH(host, user, password, wg)
		}

		if len(group.SubGroups) > 0 {
			processGroup(group.SubGroups, user, password, wg)
		}
	}
}

func runSSH(hostConfig api.HostConfig, user, password string, wg *sync.WaitGroup) {
	defer wg.Done()

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

	commands := [...]string{"ls -la", "uname -a"}

	for _, cmd := range commands {
		session, err := conn.NewSession()
		if err != nil {
			log.Printf("Error creating session for command '%s' on %s: %v", cmd, address, err)
			continue
		}
		defer session.Close()

		log.Printf("Running command '%s' on %s...", cmd, address)

		output, err := session.CombinedOutput(cmd)
		if err != nil {
			log.Printf("Error running command '%s' on %s: %v", cmd, address, err)
			continue
		}

		fmt.Printf("\n--- Output from %s for command '%s' ---\n%s\n", address, cmd, string(output))
	}

	fmt.Printf("\n--- All commands finished on %s ---\n", address)
}
