//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	head.Next.Next = swapPairs(head.Next.Next)

	help := head.Next
	head.Next = help.Next
	help.Next = head

	return help
}

func main() {
	s := []int{1, 2, 3, 4}
	head := new(ListNode)
	dummy := head
	for _, v := range s {
		head.Next = &ListNode{v, nil}
		head = head.Next
	}

	a := swapPairs(dummy.Next)

	for a != nil {
		fmt.Println(dummy)
		a = a.Next
	}
}