package BFS

import (
	"fmt"
	"testing"
)

func TestBfs(t *testing.T) {

	graph := make(Graph)
	graph[0] = []int{1, 2}
	graph[1] = []int{0, 3, 4}
	graph[2] = []int{0, 5}
	graph[3] = []int{1}
	graph[4] = []int{1, 6}
	graph[5] = []int{2}
	graph[6] = []int{4}

	source := 0
	target := 6
	fmt.Printf("Shortest path from %d to %d: %v\n", source, target, Bfs(graph, source, target))

}
