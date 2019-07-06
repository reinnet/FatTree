package topology

// Node is a physical server in fat tree topology
type Node struct {
	ID              string   `yaml:"id"`
	Cores           int      `yaml:"cores"`
	RAM             int      `yaml:"ram"`
	VNFSupport      bool     `yaml:"vnfSupport"`
	NotManagerNodes []string `yaml:"notManagerNodes"`
	Egress          bool     `yaml:"egress"`
	Ingress         bool     `yaml:"ingress"`
}
