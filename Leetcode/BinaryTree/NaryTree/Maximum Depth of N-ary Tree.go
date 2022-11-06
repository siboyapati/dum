package NaryTree

//559. Maximum Depth of N-ary Tree
type Node struct {
	Val      int
	Children []*Node
}

//// Iterative Approch
//func maxDepth(root *Node) int {
//	var queue []*Node
//	var result int
//	if root == nil {
//		return 0
//	}
//	queue = append(queue, root)
//	for len(queue) != 0 {
//		size := len(queue)
//		for i := 0; i < size; i++ {
//			temp := queue[0]
//			queue = queue[1:]
//			for k := 0; k < len(temp.Children); k++ {
//				child := temp.Children[k]
//				queue = append(queue, child)
//			}
//		}
//		result++
//	}
//	return result
//}

//func maxDepth(root *Node) int {
//	if root == nil {
//		return 0
//	}
//	return helper(root)
//}
//func helper(root *Node) int {
//	if root == nil {
//		return 0
//	}
//	var depth int
//	for i := 0; i < len(root.Children); i++ {
//		depth = Max(depth, helper(root.Children[i]))
//	}
//	return depth + 1
//
//}
//func Max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
