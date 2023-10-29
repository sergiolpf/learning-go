package DFS

import "fmt"

type Graph map[int][]int

func Dfs(graph Graph, node int, visited map[int]bool) {
	if visited[node] {
		return
	}

	visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			Dfs(graph, neighbor, visited)
		}
	}

}
