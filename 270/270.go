//给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。
//
//注意：
//
//给定的目标值 target 是一个浮点数
//题目保证在该二叉搜索树中只会存在一个最接近目标值的数
//示例：
//
//输入: root = [4,2,5,1,3]，目标值 target = 3.714286
//
//4
/// \
//2   5
/// \
//1   3
//
//输出: 4

package main

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var small, big int


// 中序遍历，找出小于和大于target最近的节点
func closestValue(root *TreeNode, target float64) int {
	small, big = math.MinInt64, math.MaxInt64
	b := help(root, target)

	if b == false {
		return small
	}
	if  target - float64(small) > float64(big) - target {
		return big
	}

	return small
}

func help(root *TreeNode, target float64) bool {
	if root == nil {
		return false
	}

	if help(root.Left, target) == true {
		return true
	}

	if float64(root.Val) > target {
		big = root.Val
		return true
	}

	small = root.Val

	if help(root.Right, target) == true {
		return true
	}

	return false
}
