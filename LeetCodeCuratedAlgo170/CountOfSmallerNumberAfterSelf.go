package main

import (
	"container/heap"
	"fmt"
)

//Input: nums = [5,2,6,1]
//Output: [2,1,1,0]
func countSmaller(nums []int) []int {
	var result []int
	h := &maxHeap{}
	heap.Init(h)
	for i := len(nums) - 1; i >= 0; i-- {
		index := 0
		cur := nums[i]
		if h.Len() == 0 {
			result = append(result, 0)
			heap.Push(h, cur)
			continue
		}

		for index < h.Len() && cur <= (*h)[index] {
			index++
		}
		result = append(result, h.Len()-index)
		heap.Push(h, cur)
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	fmt.Println(result)
	return result
}

//implementation of maxHeap
type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func main() {
	//nums := []int{5, 2, 6, 1}
	//nums := []int{-1}
	nums := []int{0, 1, 2}
	fmt.Println(countSmaller(nums))
}
