//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
//示例 1:
//
//输入: n = 12
//输出: 3
//解释: 12 = 4 + 4 + 4.
//示例 2:
//
//输入: n = 13
//输出: 2
//解释: 13 = 4 + 9.

package main

import "math"

// 转移方程：dp[i] = min(dp[i], dp[i - j * j] + 1)

func numSquares(n int) int {
	var dp = make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt16
		for j := int(math.Sqrt(float64(i))); j > 0; j-- {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-j*j])+1))
		}
	}

	return dp[n]
}

func main() {

}
