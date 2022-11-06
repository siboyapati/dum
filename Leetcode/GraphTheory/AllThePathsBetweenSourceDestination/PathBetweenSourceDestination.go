package main

import "fmt"

func allPathsSourceTarget_dfs(graph [][]int) [][]int {

	result := make([][]int, 0)
	q := [][]int{{0}}

	for len(q) != 0 {
		temp := q[0]
		q = q[1:]
		lastIndex := temp[len(temp)-1]
		if lastIndex == len(graph)-1 {
			result = append(result, temp)
		} else {
			for _, v := range graph[lastIndex] {
				newTemp := make([]int, len(temp))
				copy(newTemp, temp)
				newTemp = append(newTemp, v)
				q = append(q, newTemp)
			}
		}
	}
	return result
}

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	temp := []int{0}
	helper(graph, &result, temp, 0)
	return result
}

func helper(graph [][]int, result *[][]int, temp []int, index int) {
	if index == len(graph)-1 {
		*result = append(*result, append([]int{}, temp...))
	}
	for _, v := range graph[index] {
		temp = append(temp, v)
		helper(graph, result, temp, v)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	arr := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget_dfs(arr))
}

//
//func allPathsSourceTarget_bfs2(graph [][]int) [][]int {
//
//	q := make([][]int, 0)
//	adj := make([])
//	for _, g := range graph {
//
//	}
//
//}
