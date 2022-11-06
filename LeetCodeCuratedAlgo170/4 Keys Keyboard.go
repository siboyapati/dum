package main

import (
	"fmt"
	"sort"
)

//Input: items = [[1,46],[1,93],[1,94],[1,36],[1,71],[2,4],[2,68],[2,63],[2,80],[2,27],[3,79],[3,35],[3,95],[3,56],[3,35],[4,31],[4,32],[4,82],[4,42],[4,97]]
//Output: [[1,87],[2,88]]

func highFive(items [][]int) [][]int {

	var result [][]int
	m := make(map[int][]int)
	for _, item := range items {
		m[item[0]] = append(m[item[0]], item[1])
	}
	fmt.Println("m", m)
	for k, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
		var sum int
		for i := 0; i < 5; i++ {
			sum += v[i]
		}
		result = append(result, []int{k, sum / 5})

	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

func main() {
	items := [][]int{{1, 46}, {1, 93}, {1, 94}, {1, 36}, {1, 71}, {2, 4}, {2, 68}, {2, 63}, {2, 80}, {2, 27}, {3, 79}, {3, 35}, {3, 95}, {3, 56}, {3, 35}, {4, 31}, {4, 32}, {4, 82}, {4, 42}, {4, 97}}
	fmt.Println(highFive(items))
}
