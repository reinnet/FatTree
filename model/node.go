package model

// Node is a physical server in the topology.
type Node struct {
	Cores           int      `yaml:"cores"`
	RAM             int      `yaml:"ram"`
	VNFSupport      bool     `yaml:"vnfSupport"`
	Egress          bool     `yaml:"egress"`
	Ingress         bool     `yaml:"ingress"`
	ID              string   `yaml:"id"`
	NotManagerNodes []string `yaml:"notManagerNodes"`
}
