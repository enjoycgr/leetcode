//给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
//
//你应当保留两个分区中每个节点的初始相对位置。
//
//示例:
//
//输入: head = 1->4->3->2->5->2, x = 3
//输出: 1->2->2->4->3->5

package main

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	big, small :=  new(ListNode), new(ListNode)
	dummySmall, dummyBig := small, big
	for head != nil {
		if head.Val < x {
			small.Next = &ListNode{head.Val, nil}
			small = small.Next
		} else {
			big.Next = &ListNode{head.Val, nil}
			big = big.Next
		}
		head = head.Next
	}
	small.Next = dummyBig.Next

	return dummySmall.Next
}

func main() {
	s := []int{1, 4, 3, 2, 5, 2}
	head := new(ListNode)
	dummy := head
	for _, v := range s {
		head.Next = &ListNode{v, nil}
		head = head.Next
	}

	partition(dummy.Next, 3)
}