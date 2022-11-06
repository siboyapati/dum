package CloneGraph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	m := make(map[int]*Node)
	visited := make(map[int]bool)

	q := make([]*Node, 0)
	q = append(q, node)

	for len(q) > 0 {
		temp := q[0]
		q = q[1:]
		m[temp.Val] = &Node{
			Val: temp.Val,
		}
		visited[temp.Val] = true
		for _, v := range temp.Neighbors {
			if !visited[v.Val] {
				q = append(q, v)
			}
		}
	}

	q1 := make([]*Node, 0)
	q1 = append(q, node)
	visited1 := make(map[int]bool)

	for len(q1) > 0 {
		temp := q1[0]
		q1 = q1[1:]
		visited[temp.Val] = true
		for _, v := range temp.Neighbors {
			temp.Neighbors = append(temp.Neighbors, m[v.Val])
			if !visited1[v.Val] {
				q1 = append(q1, v)
			}
		}
	}
	return m[node.Val]
}
