package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {

	inDegree := make([]int, numCourses)
	m := make(map[int][]int)
	q := make([]int, 0)
	var result []int

	for i := 0; i < len(prerequisites); i++ {
		inDegree[prerequisites[i][0]]++
		m[prerequisites[i][1]] = append(m[prerequisites[i][1]], prerequisites[i][0])
	}
	fmt.Println("map", m)
	fmt.Println(inDegree)

	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) != 0 {
		temp := q[0]
		q = q[1:]
		result = append(result, temp)
		for _, val := range m[temp] {
			inDegree[val]--
			if inDegree[val] == 0 {
				q = append(q, val)
			}
		}

	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}

func main() {
	//arr := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	//numCourses := 4
	//
	//fmt.Println("final result ###", findOrder(numCourses, arr))
	//m := make(map[int][]int)
	var m map[int][]int
	m[1] = append(m[1], 1, 2, 3)
	fmt.Println(m)
}

//todo: check
//Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
//Output: [0,2,1,3]
func findOrder_bfs(numCourses int, prerequisites [][]int) []int {

	var m map[int][]int
	//var result []int
	var tempArr []int
	visited := make([]int, numCourses)
	for _, pre := range prerequisites {
		curC, preC := pre[0], pre[1]
		m[preC] = append(m[preC], curC)
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] != 1 && dfs(visited, m, i, &tempArr) {
			return []int{}
		}
	}
	return tempArr
}

//
//0-1,2
//1-3
//2-3

func dfs(visited []int, m map[int][]int, i int, tempArr *[]int) bool {

	//visited[i] = -1
	if visited[i] == -1 {
		return false
	}
	visited[i] = -1
	for _, v := range m[i] {
		if visited[v] != 1 {
			dfs(visited, m, v, tempArr)
		}
	}
	visited[i] = 1
	*tempArr = append([]int{1}, *tempArr...)
	return true
}
