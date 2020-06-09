//计算给定二叉树的所有左叶子之和。
//
//示例：
//
//3
/// \
//9  20
///  \
//15   7
//
//在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var s int

func sumOfLeftLeaves(root *TreeNode) int {
	s = 0
	sums(root)
	return s
}

func sums(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if sums(root.Left) == true {
		s = s + root.Left.Val
	}

	sums(root.Right)

	return false
}
