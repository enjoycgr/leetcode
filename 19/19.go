//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//说明：
//
//给定的 n 保证是有效的。


package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

// 快慢指针，两指针中间隔着n个， first为nil的时候，second.Next为要删除的node
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{0, head}
	first, second := h, h
	i := 0

	for first != nil {
		first = first.Next
		if i > n {
			second = second.Next
		}
		i++
	}

	second.Next = second.Next.Next

	return h.Next
}

func main() {
	fmt.Println(removeNthFromEnd(&ListNode{1, nil}, 1))
}
