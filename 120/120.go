//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
//
//
//例如，给定三角形：
//
//[
//   [2],
//  [3,4],
// [6,5,7],
//[4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

package main

import (
	"fmt"
	"math"
)

// 1.转移方程dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
// 2.边界，当 j = 0时, dp[i][j] = dp[i-1][j] + triangle[i][j]
// 		  当 i = 0，j = 0时，dp[0][0] = triangle[0][0]
//		  当 i = j时，dp[i][j] = dp[i-1][j-1] + triangle[i][j]
// 3.计算顺序，dp[0][0],dp[1][0],dp[1][1]...dp[len][len]

func minimumTotal(triangle [][]int) int {
	var length = len(triangle)
	var dp = make([][]int, length)
	var small = math.MaxInt32
	if length == 0 {
		return 0
	}
	if length == 1 {
		return triangle[0][0]
	}
	dp[0] = make([]int, 1)
	dp[0][0] = triangle[0][0]

	for i := 1; i < length; i++ {
		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if dp[i][j] == 0 {
				dp[i][j] = math.MaxInt32
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else {
				if j == i {
					dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				} else {
					dp[i][j] = int(math.Min(float64(dp[i-1][j-1]), float64(dp[i-1][j]))) + triangle[i][j]
				}
			}

			if dp[i][j] < small && i == length-1 {
				small = dp[i][j]
			}
		}
	}
	return small
}

func main() {
	var data = [][]int{
		{-1},
		{2, 3},
		{1, -1, -3},
	}
	fmt.Println(minimumTotal(data))
}
