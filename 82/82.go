//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//
//示例 1:
//
//输入: 1->2->3->3->4->4->5
//输出: 1->2->5
//示例 2:
//
//输入: 1->1->1->2->3
//输出: 2->3

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	flag := false
	for head.Next != nil && head.Val == head.Next.Val {
		flag = true
		head = head.Next
	}

	head.Next = deleteDuplicates(head.Next)

	if flag == true {
		return head.Next
	}

	return head
}
