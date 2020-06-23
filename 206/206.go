//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return p
}

// 迭代
func reverseList2(head *ListNode) *ListNode {
	var dummy, next *ListNode
	for head != nil {
		next = head.Next

		head.Next = dummy
		dummy = head
		head = next
	}

	return dummy
}

func main() {
	s := []int{1, 2, 3, 4}
	head := new(ListNode)
	dummy := head
	for _, v := range s {
		head.Next = &ListNode{v, nil}
		head = head.Next
	}

	a := reverseList(dummy.Next)

	for a != nil {
		fmt.Println(dummy)
		a = a.Next
	}
}
