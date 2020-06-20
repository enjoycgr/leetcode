package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	r := sortList(slow.Next)
	slow.Next = nil
	l := sortList(head)

	res := &ListNode{0, nil}
	help(l, r, res)
	return res.Next
}

func help(l, r, res *ListNode) {
	p := res

	for l != nil && r != nil {
		if l.Val > r.Val {
			p.Next = r
			r = r.Next
		} else {
			p.Next = l
			l = l.Next
		}
		p = p.Next
	}

	if l == nil {
		p.Next = r
	}

	if r == nil {
		p.Next = l
	}
}
