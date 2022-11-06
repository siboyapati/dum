package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 1, 3, 50, 75}
	lower := 0
	upper := 99
	fmt.Println(findMissingRanges(nums, lower, upper))
}

func helper(result *[]string, lower, higher int) {
	if lower == higher {
		str := strconv.Itoa(lower)
		*result = append(*result, str)
	} else if lower < higher {
		str := strconv.Itoa(lower) + "->" + strconv.Itoa(higher)
		*result = append(*result, str)
	}
}

func findMissingRanges(nums []int, lower int, upper int) []string {

	result := make([]string, 0)
	start := lower
	if len(nums) == 0 {
		helper(&result, lower, upper)
	}
	for _, num := range nums {
		if start < num {
			helper(&result, start, num-1)
		}
		start = num + 1
	}
	helper(&result, start, upper)
	return result
}
