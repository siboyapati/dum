package CycleInGraph

func union(parent []int, i int, k int) {
	iset := find(parent, i)
	kset := find(parent, k)
	if iset != kset {
		parent[iset] = kset
	}
}

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}

//public int find(int x) {
//if (x == root[x]) {
//return x;
//}
//return root[x] = find(root[x]);
//}
