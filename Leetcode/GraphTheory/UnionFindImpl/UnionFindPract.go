package main

type UnionFind struct {
	n      int
	parent []int
	count  int
}

func NewUnion(n int) UnionFind {
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return UnionFind{
		n,
		arr,
		n,
	}
}

func (this *UnionFind) Union(x, y int) {
	temp1 := this.Find(x)
	temp2 := this.Find(y)
	if temp1 != temp2 {
		this.parent[temp1] = temp2
		this.count--
	}
}

func (this *UnionFind) Find(id int) int {

	if id == this.parent[id] {
		return id
	}
	this.parent[id] = this.Find(this.parent[id])
	return this.parent[id]
}
