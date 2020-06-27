//给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。
//
//每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
//
//示例 1:
//
//输入:
//
//1
///   \
//3     2
/// \     \
//5   3     9
//
//输出: 4
//解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
//示例 2:
//
//输入:
//
//1
///
//3
/// \
//5   3
//
//输出: 2
//解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var s []int
var res int

// s记录每层的第一个节点的pos，res记录最大宽度

func widthOfBinaryTree(root *TreeNode) int {
	s = make([]int, 0)
	res = 1
	help(root, 0, 1)

	return res
}

func help(root * TreeNode, depth, pos int) {
	if root == nil {
		return
	}

	if len(s) == depth {
		s = append(s, pos)
	} else {
		if pos - s[depth] + 1 > res {
			res = pos - s[depth] + 1
		}
	}

	help(root.Left, depth + 1, pos * 2)
	help(root.Right, depth + 1, pos * 2 + 1)

	return
}
