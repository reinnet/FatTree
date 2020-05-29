package fattree

import (
	"errors"
	"math/rand"

	"github.com/reinnet/topology/fattree/namer"
	"github.com/reinnet/topology/model"
)

// FatTree builds a FatTree topology for given K.
type FatTree struct {
	k int
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

// ErrInvalidK is an error for setting an invalid value for k.
var ErrInvalidK = errors.New("given k is not a valid k for k-fattree")

// New creates a new FatTree builder and validate configurations.
func New(k int) (FatTree, error) {
	if k%2 == 1 {
		return FatTree{}, ErrInvalidK
	}

	return FatTree{
		k: k,
	}, nil
}

// nolint: gomnd, funlen
func (f FatTree) perPod(pod int) ([]model.Node, []model.Link) {
	pods := f.k

	aggregations := f.k / 2

	edges := f.k / 2

	servers := f.k * f.k / 4

	var nodes []model.Node

	var links []model.Link

	for j := 0; j < aggregations; j++ {
		nodes = append(nodes, model.Node{
			ID:              namer.AggrSwitch(pod, j),
			Cores:           0,
			RAM:             0,
			VNFSupport:      false,
			NotManagerNodes: []string{},
			Egress:          false,
			Ingress:         false,
		})
	}

	for j := 0; j < edges; j++ {
		nodes = append(nodes, model.Node{
			ID:              namer.EdgeSwitch(pod, j),
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
			links = append(links, model.Link{
				Source:      namer.AggrSwitch(pod, j),
				Destination: namer.EdgeSwitch(pod, k),
				Bandwidth:   LinkBandwidth,
			})

			links = append(links, model.Link{
				Destination: namer.AggrSwitch(pod, j),
				Source:      namer.EdgeSwitch(pod, k),
				Bandwidth:   LinkBandwidth,
			})
		}
	}

	// servers in far pods cannot manage these server
	// so create a list of them and set as not manager nodes
	var nmn []string

	for k := 0; k < pods; k++ {
		if k != pod && k != pod+1 && k != pod-1 {
			for j := 0; j < servers; j++ {
				nmn = append(nmn, namer.Server(k, j))
			}
		}
	}

	for j := 0; j < servers; j++ {
		nodes = append(nodes, model.Node{
			ID:              namer.Server(pod, j),
			Cores:           rand.Intn(CoresUB-CoresLB) + CoresLB,
			RAM:             rand.Intn(MemoryUB-MemoryLB) + MemoryLB,
			VNFSupport:      true,
			NotManagerNodes: nmn,
			Egress:          false,
			Ingress:         false,
		})

		links = append(links, model.Link{
			Source:      namer.EdgeSwitch(pod, j/(pods/2)),
			Destination: namer.Server(pod, j),
			Bandwidth:   LinkBandwidth,
		})

		links = append(links, model.Link{
			Destination: namer.EdgeSwitch(pod, j/(pods/2)),
			Source:      namer.Server(pod, j),
			Bandwidth:   LinkBandwidth,
		})
	}

	return nodes, links
}

// Build builds a k-FatTree Topology.
// nolint: gomnd
func (f FatTree) Build() model.Config {
	var nodes []model.Node

	var links []model.Link

	// globals
	pods := f.k

	cores := f.k * f.k / 4

	for i := 0; i < cores; i++ {
		nodes = append(nodes, model.Node{
			ID:              namer.CoreSwitch(i),
			Cores:           0,
			RAM:             0,
			VNFSupport:      true,
			NotManagerNodes: []string{},
			Egress:          true,
			Ingress:         true,
		})

		for j := 0; j < pods; j++ {
			links = append(links, model.Link{
				Source:      namer.CoreSwitch(i),
				Destination: namer.AggrSwitch(j, i/(pods/2)),
				Bandwidth:   LinkBandwidth,
			})

			links = append(links, model.Link{
				Destination: namer.CoreSwitch(i),
				Source:      namer.AggrSwitch(j, i/(pods/2)),
				Bandwidth:   LinkBandwidth,
			})
		}
	}

	for i := 0; i < pods; i++ {
		ns, ls := f.perPod(i)

		nodes = append(nodes, ns...)
		links = append(links, ls...)
	}

	return model.Config{
		Nodes: nodes,
		Links: links,
	}
}
