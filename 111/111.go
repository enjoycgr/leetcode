//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回它的最小深度  2.

package main

// 递归遍历左右子树
// 如果左子树为nil，那么父节点的深度=右子树 + 1
// 如果右子树为nil，那么父节点的深度=左子树 + 1
// 如果左右子树都不为nil，那么父节点的深度=min(左，右) + 1

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if root.Left == nil {
		return right + 1
	}
	if root.Right == nil {
		return left + 1
	}

	return int(math.Min(float64(left), float64(right))) + 1
}
