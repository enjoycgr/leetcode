package main

import (
	"fmt"
	"strings"
)

func main() {
	l := lengthOfLongestSubstring("bwf")
	fmt.Println(l)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var local, l int
	length := 1
	for {
		l = lengthOfString(s, s[local:], 1)
		if length < l {
			length = l
		}
		local++
		if local == len(s) {
			return length
		}
	}
	return l
}

func lengthOfString(S string, s string, i int) int {
	if i >= len(s) || strings.Index(S[:i], s[i:i+1]) != -1 {
		return l
	}
	return lengthOfString(S, s[i:], i)
}
