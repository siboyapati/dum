package main

import (
	"container/heap"
	"fmt"
)

func validTree(n int, edges [][]int) bool {

	adjMap := make(map[int][]int)
	inDegreeArr := make([]int, n)
	var q []int
	var count int

	for _, edge := range edges {
		src, dest := edge[0], edge[1]
		adjMap[src] = append(adjMap[src], dest)
		adjMap[dest] = append(adjMap[dest], src)
		inDegreeArr[src]++
		inDegreeArr[dest]++
	}

	for i, in := range inDegreeArr {
		if in <= 1 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		temp := q[0]
		q = q[1:]
		count++
		for _, v := range adjMap[temp] {
			inDegreeArr[v]--
			if inDegreeArr[v] == 1 {
				q = append(q, v)
			}
		}
	}

	if count != n {
		return false
	}
	return true
}

// function print hello world
func printHelloWorld() {
	fmt.Println("Hello World")
}

// min heap implementation
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &MinHeap{2, 1, 5, 3}
	heap.Init(h)
	heap.Push(h, 4)
	fmt.Println(*h)
	fmt.Println(heap.Pop(h))
	fmt.Println(*h)
	// array of popular charAtFirst
	arr := []string{"toyota", "honda", "bmw"}
	fmt.Println(arr)
}
