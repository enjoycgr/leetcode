package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true

func isValid(s string) bool {
	l := ""
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		if len(l) > 0 && ((l[len(l)-1] == '}' && s[i] == '{') || (l[len(l)-1] == ')' && s[i] == '(') || (l[len(l)-1] == ']' && s[i] == '[')) {
			s = s[:len(s)-1]
			l = l[:len(l)-1]
		} else {
			l = l + string(s[i])
			s = s[:len(s)-1]
		}
	}
	if len(l) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("({})"))
}
