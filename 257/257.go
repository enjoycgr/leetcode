//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//输入:
//
//1
///   \
//2     3
//\
//5
//
//输出: ["1->2->5", "1->3"]
//
//解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = make([]string, 0)
	path(root, "")
	return res
}

func path(root *TreeNode, p string) {
	if root == nil {
		return
	}

	if p != "" {
		p = p + "->"
	}

	p = p + strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		res = append(res, p)
		return
	}

	path(root.Left, p)
	path(root.Right, p)
}
