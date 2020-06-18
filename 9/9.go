//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//示例 1:
//
//输入: 121
//输出: true
//示例 2:
//
//输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//示例 3:
//
//输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。

package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10))
}

// 将数字切一半，后一半翻转 判断两部分是否相等
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	s := 0
	for x > s {
		s = s*10 + x%10
		x = x / 10
	}

	// 数组长度为奇数的话需要除以10
	return s == x || s/10 == x
}