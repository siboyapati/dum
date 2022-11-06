package main

import (
	"container/heap"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestKValues(root *TreeNode, target float64, k int) []int {

	var result []int
	var helper func(*TreeNode, float64, int)
	pq := &PriorityQueue{}
	heap.Init(pq)
	helper = func(root *TreeNode, target float64, k int) {
		if root == nil {
			return
		}
		if float64(root.Val) <= target {
			heap.Push(pq, Item{root.Val, math.Abs(target - float64(root.Val))})
		} else {
			heap.Push(pq, Item{root.Val, math.Abs(target - float64(root.Val))})
		}
		if pq.Len() > k {
			heap.Pop(pq)
		}
		helper(root.Left, target, k)
		helper(root.Right, target, k)
	}
	helper(root, target, k)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(pq).(Item).value)
	}
	return result
}

type Item struct {
	value    int
	priority float64
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Peek() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	return item
}

func (pq *PriorityQueue) Update(i int, v Item) {
	old := *pq
	n := len(old)
	item := old[i]
	item.priority = v.priority
	*pq = old[0 : n-1]
	*pq = append(*pq, item)
}
