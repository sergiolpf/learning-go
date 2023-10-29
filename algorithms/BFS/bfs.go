package BFS

import (
	"container/list"
	"fmt"
)

type Graph map[int][]int

func Bfs(graph Graph, start, target int) []int {
	visited := make(map[int]bool)
	parent := make(map[int]int)
	queue := list.New()

	visited[start] = true
	queue.PushBack(start)

	for queue.Len() > 0 {
		node := queue.Front()
		queue.Remove(node)
		fmt.Printf("%d", node.Value.(int))

		currentNode := node.Value.(int)

		if currentNode == target {
			path := []int{currentNode}
			for parent[currentNode] != start {
				currentNode = parent[currentNode]
				path = append([]int{currentNode}, path...)
			}
			path = append([]int{start}, path...)

			return path
		}

		for _, neighbor := range graph[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = currentNode
				queue.PushBack(neighbor)
			}
		}
	}

	return []int{}
}
