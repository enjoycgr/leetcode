//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
//说明：
//
//拆分时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false

package main

import "fmt"

// 转移方程：dp[j] = dp[i - 1] == true && s[i - 1:j] == true
// 边界：dp[0] = true, 从dp[1]，s[0]开始，

func wordBreak(s string, wordDict []string) bool {
	word := make(map[string]bool)
	for _, v := range wordDict {
		word[v] = true
	}

	dp := make([]bool, len(s) + 1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			if dp[i - 1] == true && word[s[i - 1:j]] == true {
				dp[j] = true
			}
		}
	}

	if dp[len(s)] == true {
		return true
	}

	return false
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet","code"}))
}