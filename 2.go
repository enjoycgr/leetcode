package main

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义一个新链表 用于结果返回
	var newList = &ListNode{
		Val:  -1,
		Next: nil,
	}

	// 赋值给一个新链表 用于链表当前节点过渡
	currentList := newList

	// 定义当前两数求和的值、上一次除以10的商
	var sum, preshang int = 0, 0

	// 因为要遍历所有链表，所以需要无限循环
	for {
		// 两个链表都为空，跳出
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil { // l1为空 取l2的值
			sum = l2.Val + preshang
		} else if l2 == nil { // l2为空 取l1的值
			sum = l1.Val + preshang
		} else { // l1不为空 l2不为空 相加
			sum = l1.Val + l2.Val + preshang
		}

		// 把新节点赋值给本次节点的Next节点
		currentList.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		// 把本次节点位置锁定到下一次节点的Next节点
		currentList = currentList.Next

		// 依次往后遍历
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		// 赋值本次的商给下一次循环的上一次商变量
		preshang = sum / 10
	}

	if preshang != 0 {
		// 最后一次别忘记处理！！！
		currentList.Next = &ListNode{
			Val:  preshang,
			Next: nil,
		}
	}

	// 返回第一个节点以后的链表
	return newList.Next
}
