package topology

// Link is a physical link in fat tree topology
type Link struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
	Bandwidth   int    `yaml:"bandwidth"`
}
