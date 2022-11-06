package main

import (
	"awesomeProject/Leetcode/Strings/LongestSubstringWithOutRepChar"
	"fmt"
)

func main() {
	fmt.Println(LongestSubstringWithOutRepChar.LengthOfLongestSubstring1("tmmzuxt"))
}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {

	var result [][]int
	var stack []*Node
	stack = append(stack, root)
	result = append(result, []int{root.Val})
	for len(stack) != 0 {
		size := len(stack)
		var tempInt []int
		for i := 0; i < size; i++ {
			temp := stack[0]
			stack = stack[1:]
			for j := 0; j < len(temp.Children); j++ {
				child := temp.Children[j]
				if child == nil {
					continue
				}
				stack = append(stack, child)
				tempInt = append(tempInt, child.Val)
			}
		}
		result = append(result, tempInt)
	}
	return result
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var result int
	helper(root, &result)
	return result

}
func helper(root *Node, result *int) {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	for k := 0; k < len(root.Children); k++ {
		depth := helper(root.Children[k], result)
		max(depth, *result)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
