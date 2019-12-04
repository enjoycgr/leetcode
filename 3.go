//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
package main

import (
	"fmt"
)

func main() {
	l := lengthOfLongestSubstring("dvdf")
	fmt.Println(l)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var start, end, l, local int
	var child [256]uint8
	for {
		if end == len(s) {
			break
		}

		if child[s[end]] == 0 {
			child[s[end]] = s[end]
			local++
			end++
		} else {
			if local > l {
				l = local
			}
			local = 0
			child = [256]uint8{}
			start = start + 1
			end = start
		}

	}
	if local > l {
		l = local
	}
	return l
}
