package main

import "fmt"

//Number of Connected Components in an Undirected Graph
/*You have a graph of n nodes. You are given an integer n and an array edges where edges[i] = [ai, bi] indicates
that there is an edge between ai and bi in the graph.

Return the number of connected components in the graph.*/

/*Input: n = 5, edges = [[0,1],[1,2],[3,4]]
Output: 2

Input: n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
Output: 1*/
func countComponents(n int, edges [][]int) int {

	adjList := make(map[int][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	visited := make(map[int]bool)
	var count int
	for i := 0; i < n; i++ {
		if _, ok := visited[i]; !ok {
			hasCycle(i, -1, adjList, visited)
			count++
		}
	}
	return count
}
func hasCycle(v int, parent int, adjList map[int][]int, visited map[int]bool) {
	if _, ok := visited[v]; ok { // we've been here before -- there's a cycle
		return
	}
	visited[v] = true
	if eList, ok := adjList[v]; ok {
		for _, n := range eList {
			if n == parent { // don't try to revisit a parent
				continue
			}
			hasCycle(n, v, adjList, visited)
		}
	}

}

func main() {
	arr := [][]int{{0, 1}, {1, 2}, {3, 4}}
	fmt.Println(countComponents(5, arr))
}
