package namer

import "fmt"

// Server creates a server identification.
func Server(pod int, id int) string {
	return fmt.Sprintf("server-%d-%d", pod, id)
}
