package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var num int
var val int

func kthSmallest(root *TreeNode, k int) int {
	val = 0
	num = 0
	help(root, k)
	return val
}

func help(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	if help(root.Left, k) == true {
		return true
	}

	num++
	if num == k {
		val = root.Val
		return true
	}

	if help(root.Right, k) == true {
		return true
	}

	return false
}
