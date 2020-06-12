//给定一个二叉树，返回它的 后序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
//
//输出: [3,2,1]

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var s []int

func postorderTraversal(root *TreeNode) []int {
	s = []int{}
	s = post(root)
	return s
}

func post(root *TreeNode) []int {
	if root == nil {
		return s
	}

	post(root.Left)
	post(root.Right)
	s = append(s, root.Val)

	return s
}
