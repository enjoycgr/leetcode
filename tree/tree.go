package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历：先输出父节点，再输出左节点，最后右节点
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Println(tree.Val)
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

// 中序遍历：先输出左节点，再输出父节点，最后右节点
func InOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	InOrder(tree.Left)
	fmt.Println(tree.Val)
	InOrder(tree.Right)
}

// 后序遍历：先输出左节点，再输出右节点，最后输出父节点
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Println(tree.Val)
}
