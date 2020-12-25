package main

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i := 0
	res := ""
	long := false
	var flag bool
	for {
		flag = false
		for _, v := range strs {
			if len(v) == 0 {
				return ""
			}
			if flag == false {
				res = res + string(v[i])
				flag = true
			}
			if v[i] != res[i] {
				return res[:i]
			}
			if len(v)-1 == i {
				long = true
			}
		}

		if long {
			return res
		}
		i++
	}
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"123", "12", "121"}))
}
