//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//示例:
//
//输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

package main

import "fmt"

// 动态规划
//1.确定状态
//	dp[i]表示前i个按键的排列组合，dp[i+1]=dp[i] + phone[i+1]上所有的字母，所以dp[i+1]取决于dp[i]

//2.转移方程
// dp[i] = d[i-1]+ phone[i+1]所有字母

//3.初始条件和边界情况
//	i - 1 > 0

//4.计算顺序
//	dp[1],dp[2],dp[3]...dp[i]
func letterCombinations(digits string) []string {
	var phone = map[byte][]byte{
		50: {97, 98, 99},
		51: {100, 101, 102},
		52: {103, 104, 105},
		53: {106, 107, 108},
		54: {109, 110, 111},
		55: {112, 113, 114, 115},
		56: {116, 117, 118},
		57: {119, 120, 121, 122},
	}

	var d = []byte(digits)
	var res = make([][]string, len(d))

	if len(d) == 0 {
		return nil
	}

	for i := 0; i < len(d); i++ {

		num, ok := phone[d[i]]
		if ok {
			if i == 0 {
				res[i] = make([]string, len(num))
				for j := 0; j < len(num); j++ {
					res[i][j] = string(num[j])
				}
				continue
			}

			n := 0
			res[i] = make([]string, len(res[i-1])*len(num))
			for _, v := range res[i-1] {
				for j := 0; j < len(num); j++ {
					res[i][n] = string(append([]byte(v), num[j]))
					n++
				}
			}
		}
	}
	return res[len(d)-1]
}

func main() {
	fmt.Println(letterCombinations("234"))
}
