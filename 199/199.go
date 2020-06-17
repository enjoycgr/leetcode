//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
//示例:
//
//输入: [1,2,3,null,5,null,4]
//输出: [1, 3, 4]
//解释:
//
//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---

package main

// 遍历二叉树，level为栈的层数，获取每一层的最新节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func rightSideView(root *TreeNode) []int {
	res = make([]int, 0)
	help(root, 0)
	return res
}

func help(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(res) == level {
		res = append(res, root.Val)
	}
	res[level] = root.Val

	help(root.Left, level+1)
	help(root.Left, level+1)

	return
}
