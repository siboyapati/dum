package main

import "fmt"

func main() {
	//
	//fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
	//fmt.Println(alienOrder([]string{"z", "x"}))
	//fmt.Println(alienOrder([]string{"z", "z"}))
	//["ac","ab","zc","zb"]
	//fmt.Println(alienOrder([]string{"ac", "ab", "zc", "zb"}))
	//"acbz"
	//fmt.Println(largestPathValue("acbz", [][]int{{0, 1}, {1, 2}, {2, 3}}))
	//Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
	//Output: 3
	fmt.Println(largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}))
}

//Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
//Output: 3
//Explanation: The path 0 -> 2 -> 3 -> 4 contains 3 nodes that are colored "a" (red in the above image).
func largestPathValue(colors string, edges [][]int) int {

	var inOrder = make(map[int]int)
	var adjMap = make(map[int][]int)

	for _, edge := range edges {

		if _, ok := inOrder[edge[1]]; !ok {
			inOrder[edge[1]] = 0
		}
		if _, ok := inOrder[edge[0]]; !ok {
			inOrder[edge[0]] = 0
		}
		inOrder[edge[1]]++
		adjMap[edge[0]] = append(adjMap[edge[0]], edge[1])
	}
	var q []int
	for k, v := range inOrder {
		if v == 0 {
			q = append(q, k)
		}
	}
	if len(q) == 0 {
		return -1
	}

	var result int
	for _, node := range q {
		var size int
		dfs1(node, adjMap, inOrder, colors, &size)
		result = max(result, size)
	}
	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs1(node int, adjMap map[int][]int, inOrder map[int]int, colors string, size *int) {

	if colors[node] == 'a' {
		*size++
	}
	for _, adj := range adjMap[node] {
		dfs1(adj, adjMap, inOrder, colors, size)
	}

}