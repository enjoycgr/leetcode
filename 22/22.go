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

var res []string

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	help(n, 0, 0, []byte{})
	return res
}

func help(n, l, r int, output []byte) {
	if l + r == n * 2 {
		tmp := make([]byte, len(output))
		copy(tmp, output)
		res = append(res, string(tmp))
		return
	}

	for i := 0; i < n * 2; i++ {
		if l >= r && l < 3{
			l++
			output = append(output, '(')
			help(n, l, r, output)
			l--
		} else if r < 3 {
			output = append(output, ')')
			r++
			help(n, l, r, output)
			r--
		}
		output = output[:len(output) - 1]
	}
}

func main() {
	generateParenthesis(3)
}
