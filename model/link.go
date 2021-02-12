package model

// Link is a physical link in the topology.
type Link struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
	Bandwidth   int    `yaml:"bandwidth"`
}
