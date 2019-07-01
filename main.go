package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// Node in physical server or node in fat tree topology
type Node struct {
	ID              string   `yaml:"id"`
	Cores           int      `yaml:"cores"`
	RAM             int      `yaml:"ram"`
	VNFSupport      bool     `yaml:"vnfSupport"`
	NotManagerNodes []string `yaml:"notManagerNodes"`
}

func main() {
	var nodes []Node

	var k int

	if _, err := fmt.Scanf("%d", &k); err != nil {
		return
	}

	// globals
	pods := k
	cores := k * k / 4

	// per pods
	aggregations := k / 2
	edges := k / 2
	servers := k * k / 4

	fmt.Printf("Pods: %d\n", pods)
	fmt.Printf("Cores: %d\n", cores)
	fmt.Printf("Aggregation: %d\n", aggregations)
	fmt.Printf("Edges: %d\n", edges)
	fmt.Printf("Servers: %d\n", servers)

	fmt.Printf("Nodes: %d\n", cores+pods*(servers+edges+aggregations))

	for i := 0; i < cores; i++ {
		nodes = append(nodes, Node{
			ID:              fmt.Sprintf("switch-%d", i),
			Cores:           0,
			RAM:             0,
			VNFSupport:      false,
			NotManagerNodes: []string{},
		})
	}

	b, err := yaml.Marshal(nodes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
