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

var help *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Right)
	flatten(root.Left)

	root.Right = help
	root.Left = nil
	help = root
}
