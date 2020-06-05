package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历:根->左->右
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}
