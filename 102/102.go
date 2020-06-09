//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其层次遍历结果：
//
//[
//[3],
//[9,20],
//[15,7]
//]

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, l int) {
	if root == nil {
		return
	}

	if l == len(res) {
		res = append(res, []int{})
	}

	res[l] = append(res[l], root.Val)
	dfs(root.Left, l+1)
	dfs(root.Right, l+1)
}
