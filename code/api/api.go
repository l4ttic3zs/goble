package api

type HostConfig struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Hosts struct {
	Groups []Group `yaml:"groups"`
}

type Group struct {
	Name      string       `yaml:"name"`
	Hosts     []HostConfig `yaml:"hosts"`
	SubGroups []Group      `yaml:"groups"`
}

type Flags struct {
	User      string
	Password  string
	HostsFile string
	Commands  string
}

type Plan struct {
	Name     string `yaml:"name"`
	Hosts    string `yaml:"hosts"`
	Commands string `yaml:"commands"`
}

type Commands struct {
	Command []string `yaml:"commands"`
}
