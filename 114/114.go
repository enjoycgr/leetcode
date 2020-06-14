//给定一个二叉树，原地将它展开为一个单链表。
//例如，给定二叉树
//
//1
/// \
//2   5
/// \   \
//3   4   6
//将其展开为：
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var last *TreeNode

func flatten(root *TreeNode) {
	last = nil
	help(root)
}

func help(root *TreeNode) {
	if root == nil {
		return
	}

	help(root.Right)
	help(root.Left)

	root.Right = last
	root.Left = nil
	last = root
}
