//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
//说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
//示例:
//
//输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy, first, second := head, head, head
	i := 1
	for i < n {
		if i > n - m + 1 {
			second = second.Next
		}
		first = first.Next
		i++
	}
	next := first.Next
	last := second.Next
	second.Next = turn(second.Next, next)
	last.Next = next

	return dummy
}

func turn(head, stop *ListNode) *ListNode {
	if head == stop || head.Next == stop {
		return head
	}

	p := turn(head.Next, stop)

	head.Next.Next = head
	head.Next = nil

	return p
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	head := new(ListNode)
	dummy := head
	for _, v := range s {
		head.Next = &ListNode{v, nil}
		head = head.Next
	}

	a := reverseBetween(dummy.Next, 2, 4)

	fmt.Println(a)
}
