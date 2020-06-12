//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
//输入: [1,null,2,3]
//1
//\
//2
///
//3
//
//输出: [1,2,3]

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func preorderTraversal(root *TreeNode) []int {
	res = []int{}
	preorder(root)
	return res
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	preorder(root.Left)
	preorder(root.Right)
}
