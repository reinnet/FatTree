package namer

import "fmt"

// EdgeSwitch creates the edge switch identification
func EdgeSwitch(pod int, id int) string {
	return fmt.Sprintf("edge-switch-%d-%d", pod, id)
}

// AggrSwitch creates the aggregation switch identification
func AggrSwitch(pod int, id int) string {
	return fmt.Sprintf("aggr-switch-%d-%d", pod, id)
}

// CoreSwitch creates the core switch identification
func CoreSwitch(id int) string {
	return fmt.Sprintf("core-switch-%d", id)
}
