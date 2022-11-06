package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {

	var i int
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	if i == 0 {
		sort.Ints(nums)
		return
	}
	for j := len(nums) - 1; j >= i; j-- {
		if nums[j] > nums[i-1] {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			break
		}
	}
	sort.Ints(nums[i:])
}

func getPalindromePermutation(s string) []string {
	var result []string
	var result2 []string
	var temp []rune
	var visited []bool
	var helper func([]rune, []bool)
	var count int

	m := make(map[rune]int)
	var tempArray []rune
	var newByte rune
	for _, v := range s {
		m[v]++
	}

	for k, v := range m {
		if v%2 == 0 {
			tempArray = append(tempArray, k)
		} else if v == 3 {
			tempArray = append(tempArray, k)
			count++
			newByte = k
		} else {
			count++
			newByte = k
		}
	}

	if count > 1 {
		return result
	}
	helper = func(s []rune, visited []bool) {
		if len(temp) == len(s) {
			result = append(result, string(temp))
			return
		}
		for i := 0; i < len(s); i++ {
			if visited[i] {
				continue
			}
			if i > 0 && s[i] == s[i-1] && !visited[i-1] {
				continue
			}
			temp = append(temp, s[i])
			visited[i] = true
			helper(s, visited)
			temp = temp[:len(temp)-1]
			visited[i] = false
		}
	}
	visited = make([]bool, len(s))
	helper(tempArray, visited)

	for i := 0; i < len(result); i++ {
		if newByte == 0 {
			result2 = append(result2, result[i]+reverse(result[i]))
		} else {
			result2 = append(result2, result[i]+string(newByte)+reverse(result[i]))
		}
	}
	return result2
}

// function to reverse the string
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	s := "aaaaaa"
	fmt.Println(getPalindromePermutation(s))
}
