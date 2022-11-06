package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}
	cur := root;
	q := make([]*Node, 0)
	q = append(q, cur)

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			temp := q[0]
			q = q[1:]
			if temp.Left != nil {
				q = append(q, temp.Left)
			}
			if temp.Right != nil {
				q = append(q, temp.Right)
			}
		}

		for i := 1; i < len(q); i++ {
			q[i-1].Next = q[i]
		}
		//q[len(q)-1].Next = nil

	}
	return root
}