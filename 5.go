//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。

package main

// 动态规划
//1.确定状态
//	dp[i][j]表示子串s[i..j]是否为回文子串，这里子串s[i..j]定义为左闭右闭区间，可以取到s[i]和s[j]

//2.转移方程
//	d[i][j] = s[i] == s[j] and dp[i+1][j-1]
//	动态规划事实上是在填一张二维表格，由于构成子串，因此i和j的关系是i <= j，因此，只需要填这张表格对角线以上的部分

//3.初始条件和边界情况
//	边界情况：s[i] == s[j] 且 j - i < 3时，dp[i][j] = true，否则执行转移方程
//  初始化：单个字符一定是回文串，所以dp[i][i] = true

//4.计算顺序
//	f[0],f[1],f[2].....,f[11]

func longestPalindrome(s string) string {
	var b = []byte(s)
	var length = len(b)
	var maxLen = 0
	var start = 0

	var dp = make([][]bool, length)
	for d, _ := range dp {
		dp[d] = make([]bool, length)
	}

	if length == 1 || length == 0 {
		return s
	}

	for j := 0; j < length; j++ {
		for i := 0; i < j; i++ {
			if b[i] == b[j] && j-i < 3 {
				dp[i][j] = true
				if j-i > maxLen {
					maxLen = j - i
					start = i
				}
				continue
			}
			if b[i] == b[j] && dp[i+1][j-1] == true {
				dp[i][j] = true
				if j-i > maxLen {
					maxLen = j - i
					start = i
				}
				continue
			}
		}
	}

	return string(b[start : start+maxLen+1])
}
