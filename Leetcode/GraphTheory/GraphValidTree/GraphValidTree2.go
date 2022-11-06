package main

//func validTree(n int, edges [][]int) bool {
//
//	if len(edges) != n-1 {
//		return false
//	}
//
//	adjList := make([][int, 4)
//
//	fmt.Println(adjList)
//
//
//	for _, edge := range edges {
//		x, y := edge[0], edge[1]
//
//	}
//

func ValidTree1(n int, edges [][]int) bool {
	// You don't need these two if cases, but they'll save you some work.
	if len(edges) != n-1 {
		// since there can be no cycles, the num of edges has to
		// be one less than the num of nodes (just picture a tree)
		// if there's less, then something isn't connected ... if
		// there's more then there's a cycle
		return false
	}
	if len(edges) == 0 {
		// empty graph is a valid tree
		return true
	}

	adjList := make(map[int][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make(map[int]bool)
	hasCycle1(0, -1, adjList, visited)

	// some vertices weren't visited
	if len(visited) != n {
		return false
	}
	return true
}

func hasCycle1(v int, parent int, adjList map[int][]int, visited map[int]bool) {
	visited[v] = true
	if eList, ok := adjList[v]; ok {
		for _, n := range eList {
			if visited[n] {
				continue
			}
			hasCycle1(n, v, adjList, visited)
		}
	}
}

func main() {

	adjList := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	n := 5
	ValidTree1(n, adjList)
}
