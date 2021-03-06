//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//示例 1:
//
//输入:
//2
/// \
//1   3
//输出: true
//示例 2:
//
//输入:
//5
/// \
//1   4
/// \
//3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//根节点的值为 5 ，但是其右子节点值为 4 。

package main

// 中序遍历，二叉搜索树的上一个节点一定小于当前节点

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var help int

func isValidBST(root *TreeNode) bool {
	help = math.MinInt64
	return valid(root)
}

func valid(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if valid(root.Left) == false {
		return false
	}
	if root.Val < help {
		return false
	}
	help = root.Val
	if valid(root.Right) == false {
		return false
	}
	return true
}
