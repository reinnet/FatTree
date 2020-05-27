package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/roadtomsc/FatTree/namer"
	"github.com/roadtomsc/FatTree/topology"

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

	var k int

	if _, err := fmt.Scanf("%d", &k); err != nil {
		return
	}

	if k%2 == 1 {
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
		nodes = append(nodes, topology.Node{
			ID:              namer.CoreSwitch(i),
			Cores:           0,
			RAM:             0,
			VNFSupport:      true,
			NotManagerNodes: []string{},
			Egress:          true,
			Ingress:         true,
		})

		for j := 0; j < pods; j++ {
			links = append(links, topology.Link{
				Source:      namer.CoreSwitch(i),
				Destination: namer.AggrSwitch(j, i/(pods/2)),
				Bandwidth:   40 * 1000,
			})

			links = append(links, topology.Link{
				Destination: namer.CoreSwitch(i),
				Source:      namer.AggrSwitch(j, i/(pods/2)),
				Bandwidth:   40 * 1000,
			})
		}
	}

	for i := 0; i < pods; i++ {
		for j := 0; j < aggregations; j++ {
			nodes = append(nodes, topology.Node{
				ID:              namer.AggrSwitch(i, j),
				Cores:           0,
				RAM:             0,
				VNFSupport:      false,
				NotManagerNodes: []string{},
				Egress:          false,
				Ingress:         false,
			})
		}
		for j := 0; j < edges; j++ {
			nodes = append(nodes, topology.Node{
				ID:              namer.EdgeSwitch(i, j),
				Cores:           0,
				RAM:             0,
				VNFSupport:      false,
				NotManagerNodes: []string{},
				Egress:          false,
				Ingress:         false,
			})
		}

		for j := 0; j < aggregations; j++ {
			for k := 0; k < edges; k++ {
				links = append(links, topology.Link{
					Source:      namer.AggrSwitch(i, j),
					Destination: namer.EdgeSwitch(i, k),
					Bandwidth:   40 * 1000,
				})

				links = append(links, topology.Link{
					Destination: namer.AggrSwitch(i, j),
					Source:      namer.EdgeSwitch(i, k),
					Bandwidth:   40 * 1000,
				})
			}
		}

		// servers in far pods cannot manage these server
		// so create a list of them and set as not manager nodes
		var nmn []string
		for k := 0; k < pods; k++ {
			if k != i && k != i+1 && k != i-1 {
				for j := 0; j < servers; j++ {
					nmn = append(nmn, namer.Server(k, j))
				}
			}
		}

		for j := 0; j < servers; j++ {
			nodes = append(nodes, topology.Node{
				ID:              namer.Server(i, j),
				Cores:           rand.Intn(38) + 10,   // each server has at least 10 vCPU and 38 vCPU at most
				RAM:             rand.Intn(600) + 100, // each server has at least 100GB of ram and 700GB at most
				VNFSupport:      true,
				NotManagerNodes: nmn,
				Egress:          false,
				Ingress:         false,
			})
			links = append(links, topology.Link{
				Source:      namer.EdgeSwitch(i, j/(pods/2)),
				Destination: namer.Server(i, j),
				Bandwidth:   40 * 1000,
			})
			links = append(links, topology.Link{
				Destination: namer.EdgeSwitch(i, j/(pods/2)),
				Source:      namer.Server(i, j),
				Bandwidth:   40 * 1000,
			})
		}
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
