package DFS

import (
	"fmt"
	"testing"
)

func TestDfs(t *testing.T) {
	graph := make(Graph)
	graph[0] = []int{1, 2}
	graph[1] = []int{0, 3, 4}
	graph[2] = []int{0, 5}
	graph[3] = []int{1}
	graph[4] = []int{1, 6}
	graph[5] = []int{2}
	graph[6] = []int{4}

	visited := make(map[int]bool)

	fmt.Println("Depth-First Search starting from node 0:")
	Dfs(graph, 0, visited)
}
