package usnet

import (
	"fmt"
	"math/rand"

	"github.com/reinnet/topology/model"
)

// USNet builds an USNet topology.
type USNet struct {
}

// New creates a new USNet builder and validate configurations.
func New() (USNet, error) {
	return USNet{}, nil
}

// LinkBandwidth for each link in topology.
const LinkBandwidth = 40 * 1000

// CoresLB represents a lower bound for number of cores in each node.
const CoresLB = 10

// CoresUB represents an upper bound for number of cores in each node.
const CoresUB = 48

// MemoryLB represents a lower bound for amount of memory in each node.
const MemoryLB = 100

// MemoryUB represents an upper bound for amount of memory in each node.
const MemoryUB = 700

// NodesLB represents a lower bound for number of nodes that are attached to a point.
const NodesLB = 3

// NodesUB represents an upper bound for number of nodes that are attached to a point.
const NodesUB = 4

// Build builds a USNet Topology.
// nolint: gomnd, funlen
func (u USNet) Build() model.Config {
	var nodes []model.Node

	links := make([]model.Link, 0)

	for i := 0; i < 24; i++ {
		nodes = append(nodes, model.Node{
			ID:              fmt.Sprintf("switch-%d", i+1),
			Cores:           0,
			RAM:             0,
			VNFSupport:      true,
			NotManagerNodes: []string{},
			Egress:          rand.Intn(3) == 0,
			Ingress:         rand.Intn(3) == 0,
		})

		for j := 0; j < rand.Intn(NodesUB-NodesLB)+NodesUB; j++ {
			nodes = append(nodes, model.Node{
				ID:              fmt.Sprintf("server-%d-%d", i+1, j+1),
				Cores:           rand.Intn(CoresUB-CoresLB) + CoresLB,
				RAM:             rand.Intn(MemoryUB-MemoryLB) + MemoryLB,
				VNFSupport:      true,
				NotManagerNodes: []string{},
				Egress:          false,
				Ingress:         false,
			})

			links = append(links, model.Link{
				Source:      fmt.Sprintf("server-%d-%d", i+1, j+1),
				Destination: fmt.Sprintf("switch-%d", i+1),
				Bandwidth:   LinkBandwidth,
			})

			links = append(links, model.Link{
				Source:      fmt.Sprintf("switch-%d", i+1),
				Destination: fmt.Sprintf("server-%d-%d", i+1, j+1),
				Bandwidth:   LinkBandwidth,
			})
		}
	}

	// https://ars.els-cdn.com/content/image/1-s2.0-S0140366415003618-gr10.jpg
	connections := [][2]int{
		{1, 2},
		{1, 6},
		{2, 6},
		{2, 3},
		{3, 7},
		{3, 4},
		{3, 5},
		{4, 5},
		{4, 7},
		{5, 8},
		{6, 7},
		{6, 9},
		{6, 11},
		{7, 9},
		{7, 8},
		{8, 10},
		{9, 10},
		{9, 11},
		{9, 12},
		{10, 13},
		{10, 14},
		{11, 12},
		{11, 19},
		{11, 15},
		{12, 16},
		{12, 13},
		{13, 14},
		{13, 17},
		{14, 18},
		{15, 20},
		{15, 16},
		{16, 21},
		{16, 22},
		{16, 17},
		{17, 22},
		{17, 23},
		{17, 18},
		{18, 24},
		{19, 20},
		{20, 21},
		{21, 22},
		{22, 23},
		{23, 24},
	}

	if len(connections) != 43 {
		panic("usnet must have 43 links")
	}

	for _, c := range connections {
		if c[0] >= c[1] {
			panic("link source must be lower than its destination")
		}

		links = append(links, model.Link{
			Source:      fmt.Sprintf("switch-%d", c[0]),
			Destination: fmt.Sprintf("switch-%d", c[1]),
			Bandwidth:   LinkBandwidth,
		})

		links = append(links, model.Link{
			Source:      fmt.Sprintf("switch-%d", c[1]),
			Destination: fmt.Sprintf("switch-%d", c[0]),
			Bandwidth:   LinkBandwidth,
		})
	}

	return model.Config{
		Nodes: nodes,
		Links: links,
	}
}
