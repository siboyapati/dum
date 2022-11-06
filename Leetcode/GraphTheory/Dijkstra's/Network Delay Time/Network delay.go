package Network_Delay_Time

import "container/heap"

type Network struct {
	source int
	weight int
}

type Heap []Network

func (this Heap) Len() int           { return len(this) }
func (this Heap) Less(i, j int) bool { return this[i].weight < this[j].weight }
func (this Heap) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this *Heap) Push(n interface{}) {
	obj := *this
	obj = append(obj, n.(Network))
}

func (this *Heap) Pop() interface{} {
	obj := *this
	temp := obj[len(obj)-1]
	obj = obj[:len(obj)-1]
	return temp
}

func networkDelayTime(times [][]int, n int, k int) int {

	adjMap := make(map[int][][]int)
	visited := make(map[int]bool)
	var result int
	for _, time := range times {
		src, dest, weight := time[0], time[1], time[2]
		adjMap[src] = append(adjMap[src], []int{src, dest, weight})
	}

	network := Network{
		source: k,
		weight: 0,
	}
	h := &Heap{}
	heap.Init(h)
	heap.Push(h, network)

	for h.Len() != 0 {
		temp := heap.Pop(h).(Network)
		if visited[temp.source] {
			continue
		}
		visited[temp.source] = true
		result = Max(result, temp.weight)
		for _, adj := range adjMap[temp.source] {
			if visited[adj[0]] {
				continue
			}
			newNode := Network{
				source: adj[0],
				weight: adj[2] + temp.weight,
			}
			heap.Push(h, newNode)
		}

	}
	if len(visited) != n {
		return -1
	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
