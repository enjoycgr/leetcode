//一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
//示例 1:
//
//输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
//示例 2:
//
//输入: "226"
//输出: 3
//解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

package main

import "fmt"

// 转移方程：
// s[i] != '0'
// s[i-1]s[i] > 26时,dp[i] = dp[i-1] + dp[i-2]
// s[i-1]s[i] <= 26时,dp[i] = dp[i-1]

// s[i] == '0'
// s[i-1]s[i] <= 26, dp[i] = dp[i-2]

// 无法翻译的情况：
// s[i-1]s[i] = 0
// 以 ‘0’ 结尾的大于 26 的数字

func numDecodings(s string) int {
	pre, cur := 1, 1
	if s[0:1] == "0" {
		return 0
	}
	for i := 1; i < len(s); i++ {
		ttmp := cur
		if s[i:i+1] == "0" {
			if s[i-1:i] == "1" || s[i-1:i] == "2" {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1:i] == "1" {
			cur = cur + pre
		} else if s[i-1:i] == "2" {
			if s[i:i+1] < "7" {
				cur = cur + pre
			}
		}
		pre = ttmp
	}
	return cur
}

func main() {
	fmt.Println(numDecodings("227"))
}
