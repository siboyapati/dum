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

func validTree(n int, edges [][]int) bool {
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
	if hasCycle(0, -1, adjList, visited) {
		return false
	}

	// some vertices weren't visited
	if len(visited) != n {
		return false
	}
	return true
}

func hasCycle(v int, parent int, adjList map[int][]int, visited map[int]bool) bool {
	if _, ok := visited[v]; ok { // we've been here before -- there's a cycle
		return true
	}
	visited[v] = true

	if eList, ok := adjList[v]; ok {
		for _, n := range eList {
			if n == parent { // don't try to revisit a parent
				continue
			}
			if hasCycle(n, v, adjList, visited) {
				return true
			}
		}
	}
	return false
}
