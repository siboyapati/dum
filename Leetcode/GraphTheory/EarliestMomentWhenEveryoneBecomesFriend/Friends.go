package main

import (
	"fmt"
	"sort"
)

type UnionFind struct {
	parent []int
	n      int
	count  int
}

func NewUnion(n int) UnionFind {

	arr := make([]int, n)
	for i, _ := range arr {
		arr[i] = i
	}
	return UnionFind{
		parent: arr,
		n:      n,
		count:  n,
	}
}
func (this *UnionFind) Union(x, y int) {
	temp1 := this.Find(x)
	temp2 := this.Find(y)

	if temp1 != temp2 {
		this.count--
		this.parent[temp1] = temp2
	}
}

func (this *UnionFind) Find(id int) int {

	if id == this.parent[id] {
		return id
	}
	this.parent[id] = this.Find(this.parent[id])
	return this.parent[id]
}

func earliestAcq(logs [][]int, n int) int {

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	graph := NewUnion(n)

	for _, log := range logs {
		graph.Union(log[1], log[2])
		if graph.count == 1 {
			return log[0]
		}
	}
	return -1

}

//{{0, 2, 0},{1, 0, 1},{3, 0, 3},{4, 1, 2},{7, 3, 1}}
func main() {
	arr := [][]int{{3, 0, 3}, {0, 2, 0}, {1, 0, 1}, {3, 0, 3}, {4, 1, 2}, {7, 3, 1}}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	fmt.Println(arr)
}
