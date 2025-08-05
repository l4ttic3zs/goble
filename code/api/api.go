package api

type HostConfig struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Hosts struct {
	HostsConfigs []HostConfig `yaml:"hosts"`
}

type Flags struct {
	User      string
	Password  string
	Hostsfile string
}
