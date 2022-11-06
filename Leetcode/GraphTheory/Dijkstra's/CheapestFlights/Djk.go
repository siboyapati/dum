package CheapestFlights

import "math"

type Node struct {
	src   int
	price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	adjMap := make(map[int][][]int)
	var level int
	var result int = math.MaxInt
	for _, flight := range flights {
		src, dst, cost := flight[0], flight[1], flight[2]
		adjMap[src] = append(adjMap[src], []int{dst, cost})
	}

	helper(src, dst, k, adjMap, &result, level, 0)
	if result == math.MaxInt {
		return -1
	}
	return result
}

func helper(src, dst, k int, adjMap map[int][][]int, result *int, level, price int) {

	if level > k+1 {
		return
	}
	if src == dst && price < *result {
		*result = price
	}

	for _, edge := range adjMap[src] {
		if price+edge[1] < *result {
			helper(edge[0], dst, k, adjMap, result, level+1, price+edge[1])
		}
	}
}
