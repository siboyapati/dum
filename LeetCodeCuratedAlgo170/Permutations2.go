package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {

	var result [][]int
	temp := make([]int, 0)
	var visited []bool
	var helper func([]int, []bool)
	sort.Ints(nums)
	helper = func(nums []int, visited []bool) {

		if len(temp) == len(nums) {
			tempArray := make([]int, len(temp))
			copy(tempArray, temp)
			result = append(result, tempArray)
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			visited[i] = true
			helper(nums, visited)
			temp = temp[:len(temp)-1]
			visited[i] = false
		}
	}
	visited = make([]bool, len(nums))
	helper(nums, visited)
	return result
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

//Input: nums = [1,1,2]
//Output:
//[[1,1,2],
//[1,2,1],
//[2,1,1]]

//Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

//func permuteUnique2(nums []int) [][]int {
//
//
//}
