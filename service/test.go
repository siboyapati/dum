package service

import "fmt"

func maxNumEdgesToRemove(n int, edges [][]int) int {

	var result int

	ufAlice := NewUnionFind(n)
	ufBob := NewUnionFind(n)

	for _, edge := range edges {
		// When edge is 3, both Alice and Bob can use it
		// So we need to union both ufAlice and ufBob
		// If the union fails, it means that the two nodes are already connected
		if edge[0] == 3 {
			if !ufAlice.Union(edge[1]-1, edge[2]-1) {
				// Counting only once sice Alice and Bob are using the same edges
				result++
			}
			ufBob.Union(edge[1]-1, edge[2]-1)
		}
	}

	for _, edge := range edges {

		// When edge is 1, only Alice can use it
		if edge[0] == 1 {
			if !ufAlice.Union(edge[1]-1, edge[2]-1) {
				result++
			}
		}

		// When edge is 2, only Bob can use it
		if edge[0] == 2 {
			if !ufBob.Union(edge[1]-1, edge[2]-1) {
				result++
			}
		}
	}

	fmt.Println("Count:", ufAlice.count, ufBob.count)

	if ufAlice.count != 1 || ufBob.count != 1 {
		return -1
	}
	return result

}

// Union Find Data Structure Implementation
type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(num int) *UnionFind {
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = i
	}
	return &UnionFind{
		arr,
		num,
	}
}

func (u *UnionFind) Union(a, b int) bool {

	parentA := u.Find(a)
	parentB := u.Find(b)
	if parentA != parentB {
		u.parent[parentA] = parentB
		u.count--
		return true
	}
	return false

}

func (u *UnionFind) Find(a int) int {

	if u.parent[a] == a {
		return a
	}
	u.parent[a] = u.Find(u.parent[a])
	return u.parent[a]
}
