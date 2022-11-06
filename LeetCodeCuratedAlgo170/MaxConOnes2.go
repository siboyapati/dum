package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {

	var right, left int
	var max int
	var count int
	for right < len(nums) {
		if nums[right] == 0 {
			count++
		}
		for count == 2 {
			if nums[left] == 0 {
				count--
			}
			left++
		}
		max = Max(max, right-left+1)
		right++
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//nums = [1,0,1,1,0]
	nums := []int{0}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
