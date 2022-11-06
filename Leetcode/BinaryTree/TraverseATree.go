package BinaryTree

//Pre Order Traversal
//In order Traversal
//Post order Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	helpPreOrder(root, &result)
	return result
}

func helpPreOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	helpPreOrder(root.Left, result)
	helpPreOrder(root.Right, result)
}
func preorderTraversalIterative(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	if root != nil {
		stack = append(stack, root)
	} else {
		return result
	}
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, temp.Val)
		if temp.Right != nil {
			stack = append(stack, temp.Right)
		}

		if temp.Left != nil {
			stack = append(stack, temp.Left)
		}
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	for len(stack) != 0 || root.Left != nil {
		if root.Left != nil {
			stack = append(stack, root.Left)
			root = root.Left
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, temp.Val)
			if temp.Right != nil {
				stack = append(stack, temp.Right)
				root = temp.Right
			}
		}
	}
	return result
}
func inorderTraversalRecursive(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
	var result []int
	helperInorder(root, &result)
	return result
}

func helperInorder(root *TreeNode, result *[]int) {
	if root.Left != nil {
		helperInorder(root.Left, result)
	}
	*result = append(*result, root.Val)
	if root.Right != nil {
		helperInorder(root.Right, result)
	}
}

func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
	var result []int
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root.Left != nil {
		helper(root.Left, result)
	}
	if root.Right != nil {
		helper(root.Right, result)
	}
	*result = append(*result, root.Val)
}
