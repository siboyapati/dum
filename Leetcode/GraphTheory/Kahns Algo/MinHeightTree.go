package main

import (
	"math"
)

func findMinHeightTrees(n int, edges [][]int) []int {

	m := make(map[int][]int)
	var resultArr []int
	height := make([]int, n)
	minHeight := math.MaxInt8
	for _, edge := range edges {
		src, dest := edge[0], edge[1]
		m[src] = append(m[src], dest)
		m[dest] = append(m[dest], src)
	}

	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		h := helper(m, visited, i, 0)
		height[i] = h
		minHeight = min(minHeight, h)
	}

	for i := 0; i < len(height); i++ {
		if height[i] == minHeight {
			resultArr = append(resultArr, i)
		}
	}
	return resultArr
}

func helper(m map[int][]int, visited []bool, cur int, nodeHeight int) int {
	visited[cur] = true
	ret := nodeHeight
	for _, v := range m[cur] {
		if !visited[v] {
			ret = max(ret, helper(m, visited, v, nodeHeight+1))
		}
	}
	visited[cur] = false
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
