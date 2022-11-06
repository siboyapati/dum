package main

import "fmt"

func minimumSemesters_p(n int, relations [][]int) int {
	adjMap := make(map[int][]int)
	inDegreeArr := make([]int, n+1)
	var q []int
	var result int
	var count int

	for _, relation := range relations {
		pre, cur := relation[0], relation[1]
		adjMap[pre] = append(adjMap[pre], cur)
		inDegreeArr[cur]++
	}

	for i := 1; i < len(inDegreeArr); i++ {
		if inDegreeArr[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		result++
		size := len(q)
		for i := 0; i < size; i++ {
			temp := q[0]
			count++
			q = q[1:]
			for _, v := range adjMap[temp] {
				inDegreeArr[v]--
				if inDegreeArr[v] == 0 {
					q = append(q, v)
				}
			}
		}
	}

	if count != n {
		return -1
	}

	return result
}

func main() {
	//arr := [][]int{{1, 3}, {2, 3}}
	arr2 := [][]int{{1, 2}, {2, 3}, {3, 1}}
	fmt.Println(minimumSemesters_p(3, arr2))
}
