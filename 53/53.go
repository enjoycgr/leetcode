//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
package main

import (
	"fmt"
	"math"
)

// 1.转移方程dp[i]表示前i个最大的子序列和
// dp[i] = max(dp[i - 1] + nums[i], nums[i])
// 2.边界: i = 0
// 4.计算顺序，dp[0],dp[1]...dp[length]

func maxSubArray(nums []int) int {
	var length = len(nums)
	var dp = make([]int, length)
	var big = nums[0]

	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]+nums[i]), float64(nums[i])))

		if dp[i] > big {
			big = dp[i]
		}
	}

	return big
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
