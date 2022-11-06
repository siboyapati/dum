package main

import "fmt"

func findMinHeightTrees_bfs(n int, edges [][]int) []int {

	adjMap := make(map[int][]int)
	inDegreeArr := make([]int, n)
	var q []int
	var result []int

	for _, edge := range edges {
		src, dest := edge[0], edge[1]
		adjMap[src] = append(adjMap[src], dest)
		adjMap[dest] = append(adjMap[dest], src)
		inDegreeArr[src]++
		inDegreeArr[dest]++
	}

	for i, in := range inDegreeArr {
		if in <= 1 {
			q = append(q, i)
		}
	}

	for n > 2 {
		size := len(q)
		n = n - size
		for i := 0; i < size; i++ {
			temp := q[0]
			q = q[1:]
			for _, v := range adjMap[temp] {
				inDegreeArr[v]--
				if inDegreeArr[v] == 1 {
					q = append(q, v)
				}
			}
		}
	}

	for _, v := range q {
		result = append(result, v)
	}
	return result
}

func main() {
	arr := [][]int{{1, 0}, {1, 2}, {1, 3}}
	fmt.Println(findMinHeightTrees_bfs(4, arr))
}
