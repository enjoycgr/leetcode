package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	res = [][]int{}
	levelOrder(root, 0)
	reverse(res)
	return res
}

func levelOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(res) == level {
		res = append(res, []int{})
	}

	res[level] = append(res[level], root.Val)
	levelOrder(root.Left, level+1)
	levelOrder(root.Right, level+1)
}

func reverse(s [][]int) [][]int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
