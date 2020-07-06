//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//示例：
//
//输入：n = 3
//输出：[
//	"((()))",
//	"(()())",
//	"(())()",
//	"()(())",
//	"()()()"
//]

package main

import "fmt"

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	help(n, 0, 0, []byte{})
	return res
}

func help(n, l, r int, output []byte) {
	// 当左右括号加起来等于n * 2 返回
	if l + r == n * 2 {
		tmp := make([]byte, len(output))
		copy(tmp, output)
		res = append(res, string(output))
		return
	}

	// 左括号 < n 继续添加
	if l < n {
		l++
		output = append(output, '(')
		help(n, l, r, output)
		l--
		output = output[:len(output) - 1]
	}

	// 右括号 只要少于左括号就可以添加
	if r < l {
		output = append(output, ')')
		r++
		help(n, l, r, output)
		r--
		output = output[:len(output) - 1]
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
