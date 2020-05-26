//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。
//
//示例 1:
//
//输入: "abc"
//输出: 3
//解释: 三个回文子串: "a", "b", "c".
//示例 2:
//
//输入: "aaa"
//输出: 6
//说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".

package main

// 参考第五题
func countSubstrings(s string) int {
	var b = []byte(s)
	var length = len(b)
	var count = 0

	if length == 1 {
		return 1
	}

	if length == 0 {
		return 0
	}

	var dp = make([][]bool, length)
	for d, _ := range dp {
		dp[d] = make([]bool, length)
	}

	for j := 0; j < length; j++ {
		for i := 0; i < j; i++ {
			if b[i] == b[j] && j-i < 3 {
				dp[i][j] = true
				count++
				continue
			}
			if b[i] == b[j] && dp[i+1][j-1] == true {
				dp[i][j] = true
				count++
				continue
			}
		}
	}

	return count + length
}
