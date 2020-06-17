//将一个 二叉搜索树 就地转化为一个 已排序的双向循环链表 。
//
//对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。
//
//特别地，我们希望可以 就地 完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中最小元素的指针。

package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var head *Node
var last *Node

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	head = nil
	last = nil
	help(root)
	head.Left = last
	last.Right = head
	return head
}

func help(root *Node) {
	if root == nil {
		if head == nil {
			head = last
		}
		return
	}

	help(root.Left)
	if last != nil {
		last.Right = root
		root.Left = last
	}
	last = root
	help(root.Right)
}
