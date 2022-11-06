package CountUnivalueSubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	IsTrue bool
	Val    int
}

func countUnivalSubtrees(root *TreeNode) int {
	var count int
	helper(root, &count)
	return count
}

func helper(root *TreeNode, count *int) NodeInfo {

	if root == nil {
		return NodeInfo{true, -1}
	}
	if root.Left == nil && root.Right == nil {
		*count++
		return NodeInfo{true, root.Val}
	}

	lNode := helper(root.Left, count)
	rNode := helper(root.Right, count)

	if lNode.IsTrue != true || rNode.IsTrue != true {
		return NodeInfo{false, root.Val}
	}

	if lNode.IsTrue == true && rNode.IsTrue == true && (lNode.Val == root.Val || lNode.Val == -1) && (rNode.Val == root.Val || rNode.Val == -1) {
		*count++
		return NodeInfo{true, root.Val}
	} else {
		return NodeInfo{false, root.Val}
	}
}
