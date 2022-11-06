package main

import (
	"sort"
)

//Min Cost to Connect All Points
//https://leetcode.com/explore/learn/card/graph/621/algorithms-to-construct-minimum-spanning-tree/3857/

type point struct {
	cost        int
	origin      int
	destination int
}

func minCostConnectPoints(points [][]int) int {

	arr := make([]point, 0)
	var count int
	u := NewUnionFind(len(points))
	var result int
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			dist := abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
			p := point{
				cost:        dist,
				origin:      i,
				destination: j,
			}
			arr = append(arr, p)
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].cost < arr[j].cost {
			return true
		}
		return false
	})

	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		//if count > len(points)-1 {
		//	continue
		//}
		if !u.union(cur.origin, cur.destination) {
			continue
		}
		count++
		result += cur.cost
	}
	return result
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return x * -1
}

//func main() {
//	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
//	fmt.Println(minCostConnectPoints(points))
//}

type UnionFind struct {
	parent []int
	n      int
	count  int
}

func NewUnionFind(n int) UnionFind {
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return UnionFind{
		parent: arr,
		count:  n,
		n:      n,
	}
}

func (this UnionFind) union(a, b int) bool {
	//this.parent[this.find(a)] = this.parent[this.find(b)]
	parentA := this.find(a)
	parentB := this.find(b)
	if parentA == parentB {
		return false
	} else {
		this.parent[parentA] = parentB
	}
	return true
}

func (this UnionFind) find(v int) int {
	if this.parent[v] != v {
		this.parent[v] = this.find(this.parent[v])
	}
	return this.parent[v]
}
