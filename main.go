package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/roadtomsc/USNet/topology"

	yaml "gopkg.in/yaml.v2"
)

// Config aggrates links and nodes as a physical topology
type Config struct {
	Nodes []topology.Node
	Links []topology.Link
}

func main() {
	var nodes []topology.Node
	var links []topology.Link

	for i := 0; i < 24; i++ {
		nodes = append(nodes, topology.Node{
			ID:              fmt.Sprintf("node-%d", i+1),
			Cores:           rand.Intn(38) + 10,   // each server has at least 10 vCPU and 38 vCPU at most
			RAM:             rand.Intn(600) + 100, // each server has at least 100GB of ram and 700GB at most
			VNFSupport:      true,
			NotManagerNodes: []string{},
			Egress:          true,
			Ingress:         true,
		})
	}

	connections := [][2]int{
		[2]int{1, 2},
		[2]int{1, 6},
		[2]int{2, 6},
		[2]int{2, 3},
		[2]int{3, 7},
		[2]int{3, 4},
		[2]int{3, 5},
		[2]int{4, 5},
		[2]int{4, 7},
		[2]int{5, 8},
		[2]int{6, 7},
		[2]int{6, 9},
		[2]int{6, 11},
		[2]int{7, 9},
		[2]int{7, 8},
		[2]int{8, 10},
		[2]int{9, 10},
		[2]int{9, 11},
		[2]int{9, 12},
		[2]int{10, 13},
		[2]int{10, 14},
		[2]int{11, 12},
		[2]int{11, 19},
		[2]int{11, 15},
		[2]int{12, 16},
		[2]int{12, 13},
		[2]int{13, 14},
		[2]int{13, 17},
		[2]int{14, 18},
		[2]int{15, 20},
		[2]int{15, 16},
		[2]int{16, 21},
		[2]int{16, 22},
		[2]int{16, 17},
		[2]int{17, 22},
		[2]int{17, 23},
		[2]int{17, 18},
		[2]int{18, 24},
		[2]int{19, 20},
		[2]int{20, 21},
		[2]int{21, 22},
		[2]int{22, 23},
		[2]int{23, 24},
	}

	if len(connections) != 43 {
		panic("usnet must have 43 links")
	}

	for _, c := range connections {
		if c[0] >= c[1] {
			panic("link source must be lower than its destination")
		}
		links = append(links, topology.Link{
			Source:      fmt.Sprintf("node-%d", c[0]),
			Destination: fmt.Sprintf("node-%d", c[1]),
			Bandwidth:   40 * 1000,
		})
	}

	b, err := yaml.Marshal(Config{
		Nodes: nodes,
		Links: links,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Create("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := f.Write(b); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.Close(); err != nil {
		return
	}
	log.Println(string(b))
}
