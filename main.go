package main

import (
	"fmt"
)

func main() {
	var k int

	fmt.Scanf("%d", &k)

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
}
