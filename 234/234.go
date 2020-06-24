//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true

package main

import "fmt"

type ListNode struct {
   Val int
   Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	first, second := head, head
	for first != nil && first.Next != nil {
		first = first.Next.Next
		second = second.Next
	}

	second = turn(second)

	for second != nil {
		if second.Val != head.Val {
			return false
		}
		second = second.Next
		head = head.Next
	}

	return true
}

func turn(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := turn(head.Next)

	head.Next.Next = head
	head.Next = nil

	return p
}

func main() {
	s := []int{1, 2, 3,5, 5, 3, 2, 1}
	head := new(ListNode)
	dummy := head
	for _, v := range s {
		head.Next = &ListNode{v, nil}
		head = head.Next
	}

	a := isPalindrome(dummy.Next)
	fmt.Println(a)
}
