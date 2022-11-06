package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {

	adj := make([][]int, n)
	visited := make([]bool, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	fmt.Println("adj list", adj)
	return helper(n, edges, source, destination, adj, visited)
}

func validPath2_bfs(n int, edges [][]int, source int, destination int) bool {

	adj := make([][]int, n)
	visited := make([]bool, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	fmt.Println("adj list", adj)
	//return helper(n, edges, source, destination, adj, visited)
	q := make([]int, 0)
	q = append(q, source)

	for len(q) != 0 {
		temp := q[0]
		q = q[1:]
		visited[temp] = true
		if temp == destination {
			return true
		}
		for _, edge := range adj[temp] {
			if !visited[edge] {
				q = append(q, edge)
			}
		}
	}
	return false
}

func helper(n int, edges [][]int, source int, destination int, adj [][]int, visited []bool) bool {

	if source == destination {
		return true
	}
	if visited[source] {
		return false
	}
	visited[source] = true
	temp := adj[source]

	for i := 0; i < len(temp); i++ {
		nextSource := temp[i]
		if visited[nextSource] {
			continue
		}
		result := helper(n, edges, nextSource, destination, adj, visited)
		if result {
			return result
		}
	}
	return false
}

func main() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source := 0
	destination := 2
	fmt.Println(validPath(n, edges, source, destination))
}

type union struct {
	n      int
	count  int
	parent []int
}

func NewUnion(n int) union {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return union{
		n:      n,
		count:  n,
		parent: arr,
	}
}

func (this *union) find(i int) int {
	if i == this.parent[i] {
		return i
	}
	this.parent[i] = this.find(this.parent[i])
	return this.parent[i]
}

func (this *union) union(one, two int) {

	oneParent := this.find(one)
	twoParent := this.find(two)
	if oneParent != twoParent {
		this.count--
		this.parent[oneParent] = twoParent
	}
}
func validPath_unionfind(n int, edges [][]int, source int, destination int) bool {
	adj := make([][]int, n)
	unionObj := NewUnion(n)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		unionObj.union(edge[0], edge[1])
	}
	fmt.Println(unionObj.count)
	return unionObj.count == 1

}
