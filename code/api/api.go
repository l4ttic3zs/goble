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

type Plan struct {
	Name     string `yaml:"name"`
	Hosts    string `yaml:"hosts"`
	Commands string `yaml:"commands"`
}
