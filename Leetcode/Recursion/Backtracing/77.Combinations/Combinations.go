package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	var result [][]int
	var tempArray []int
	helper(n, k, &result, tempArray, 1)
	return result
}

func helper(n, k int, result *[][]int, tempArray []int, index int) {

	if len(tempArray) >= k {
		tempArrayCopy := make([]int, len(tempArray))
		copy(tempArrayCopy, tempArray)
		*result = append(*result, tempArrayCopy)
		return
	}
	for i := index; i <= n; i++ {
		tempArray = append(tempArray, i)
		helper(n, k, result, tempArray, i+1)
		tempArray = tempArray[:len(tempArray)-1]
	}
}
