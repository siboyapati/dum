package ConstructABinaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
//Output: [3,9,20,null,null,15,7]
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	if inorder == nil || postorder == nil {
		return nil
	}

	if len(inorder) != len(postorder) {
		return nil
	}

	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}

	return helper(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, m)
}

func helper(inorder []int, imin, imax int, postorder []int, pmin, pmax int, m map[int]int) *TreeNode {
	if pmin > pmax {
		return nil
	}
	root := &TreeNode{
		Val: postorder[pmax],
	}
	target := postorder[pmax]
	mid := m[target]
	root.Left = helper(inorder, imin, mid-1, postorder, pmin, pmin+mid-imin-1, m)
	root.Right = helper(inorder, mid+1, imax, postorder, pmin+mid-imin, pmax-1, m)
	return root
}

//Construct Binary Tree from Preorder and Inorder Traversal
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if inorder == nil || preorder == nil {
		return nil
	}

	if len(inorder) != len(preorder) {
		return nil
	}

	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	return helper1(inorder, 0, len(inorder)-1, preorder, 0, len(preorder)-1, m)
}

func helper1(inorder []int, imin, imax int, preorder []int, pmin, pmax int, m map[int]int) *TreeNode {
	if imin > imax {
		return nil
	}

	node := &TreeNode{
		Val: preorder[pmin],
	}
	target := preorder[pmin]
	mid := m[target]
	node.Left = helper1(inorder, imin, mid-1, preorder, pmin+1, pmin+mid-imin, m)
	node.Right = helper1(inorder, mid+1, imax, preorder, pmin+mid-imin+1, pmax, m)
	return node
}
