//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
//示例:
//
//输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6

package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 分治，然后合并
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}
	return mergeLists(lists, 0, len(lists) - 1)
}

func mergeLists(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}

	mid := (r - l)/2 + l

	leftList := mergeLists(lists, l, mid)
	rightList := mergeLists(lists, mid+1, r)

	res := mergeTwoLists(leftList, rightList)

	return res
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}