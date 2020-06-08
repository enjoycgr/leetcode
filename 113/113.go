//给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//5
/// \
//4   8
///   / \
//11  13  4
///  \    / \
//7    2  5   1
//返回:
//
//[
//[5,4,11,2],
//[5,8,4,5]
//]

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	res = [][]int{}
	path := []int{}
	check(root, sum, 0, path)
	return res
}

func check(root *TreeNode, sum int, total int, path []int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	total = root.Val + total

	if root.Left == nil && root.Right == nil && total == sum {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}

	check(root.Left, sum, total, path)
	check(root.Right, sum, total, path)
}
