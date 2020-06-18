//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4


package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	list := head
	for l1 != nil || l2 != nil{
		if l1 == nil {
			list.Next = l2
			break
		}

		if l2 == nil {
			list.Next = l1
			break
		}

		if l1.Val > l2.Val {
			list.Next = l2
			l2 = l2.Next
		} else {
			list.Next = l1
			l1 = l1.Next
		}

		list = list.Next
	}

	return head.Next
}

// 递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

}

func main() {
	
}
