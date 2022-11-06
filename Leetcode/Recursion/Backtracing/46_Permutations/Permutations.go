package main

import "fmt"

//Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
//Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

func main() {
	fmt.Println("hello world")
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var result [][]int
	helper(nums, &result, 0)
	return result
}

func helper(nums []int, result *[][]int, index int) {

	if index == len(nums) {
		tempCopy := make([]int, len(nums))
		copy(tempCopy, nums)
		*result = append(*result, tempCopy)
		return
	}

	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		helper(nums, result, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
