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
	dummy := &ListNode{0, head}
	// 前后双指针
	// second指向第m-1个， first指向第n个
	first, second := dummy, dummy
	i := 0
	for i < n {
		if i > n - m {
			second = second.Next
		}
		first = first.Next
		i++
	}
	// 指向n+1个
	next := first.Next
	// 指向第m个
	last := second.Next
	// 反转链表
	second.Next = turn(last, next)
	last.Next = next

	return dummy.Next
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

	reverseBetween(dummy.Next, 1, 5)

	for dummy != nil {
		fmt.Println(dummy)
		dummy = dummy.Next
	}
}
