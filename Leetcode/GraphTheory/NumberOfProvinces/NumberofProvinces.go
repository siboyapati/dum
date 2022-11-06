package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	arr := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(arr)
	fmt.Println(findCircleNum(arr))
}

func findCircleNum(isConnected [][]int) int {

	visited := make([]bool, len(isConnected))
	var count int
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			dfs(isConnected, visited, i)
			count++
		}
	}
	fmt.Println(visited)
	return count
}

func dfs(connected [][]int, visited []bool, i int) {
	for j := 0; j < len(connected[0]); j++ {
		if connected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(connected, visited, j)
		}
	}
}

func findCircleNum_bfs(isConnected [][]int) int {

	visited := make([]bool, len(isConnected))
	var count int
	q := make([]int, 0)
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			q = append(q, i)
			for len(q) != 0 {
				temp := q[0]
				q = q[1:]
				visited[temp] = true
				for k := 0; k < len(isConnected); k++ {
					if isConnected[temp][k] == 1 && !visited[k] {
						q = append(q, k)
					}
				}
			}
			count++
		}
	}
	return count
}

func findCircleNum_unionfind(isConnected [][]int) int {

	parent := make([]int, len(isConnected))
	var count int
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
	for i := 0; i < len(isConnected); i++ {
		for k := 0; k < len(isConnected); k++ {
			if isConnected[i][k] == 1 && i != k {
				union(parent, i, k)
			}
		}
	}
	for i := 0; i < len(parent); i++ {
		if parent[i] == -1 {
			count++
		}
	}
	return count
}

func union(parent []int, i int, k int) {
	iset := find(parent, i)
	kset := find(parent, k)
	if iset != kset {
		parent[iset] = kset
	}
}

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}
