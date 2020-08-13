//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//示例 1:
//
//输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//
//输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

package main

import (
	"fmt"
)

// 转移方程:min[i] = min(min[i-1] * nums[i], nums[i], max[i-1] * )* nums[i]
//          max[i] = max(max[i-1] * nums[i], nums[i], min[i-1] * nums[i])
// 			dp = max(dp, max[i])

func maxProduct(nums []int) int {
	dp := nums[0]
	maxdp := make([]int, len(nums))
	mindp := make([]int, len(nums))

	if len(nums) == 1 {
		return nums[0]
	}
	maxdp[0] = nums[0]
	mindp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		maxdp[i] = max(maxdp[i-1]*nums[i], max(nums[i], mindp[i-1]*nums[i]))
		mindp[i] = min(maxdp[i-1]*nums[i], min(nums[i], mindp[i-1]*nums[i]))

		if maxdp[i] > dp {
			dp = maxdp[i]
		}
	}

	return dp
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxProduct([]int{4, 6, 0, -2}))
}
