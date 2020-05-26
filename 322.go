//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

//输入: coins = [1, 2, 5], amount = 11
//输出: 3
//解释: 11 = 5 + 5 + 1

package main

import (
	"math"
)

//递归
//转移方程：f[X]=min{f[X-1]+1, f[X-2]+1, f[X-5]+1}
func recursive(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := math.MaxInt32

	for _, v := range coins {
		if amount >= v {
			res = int(math.Min(float64(recursive(coins, amount-v)+1), float64(res)))
		}
	}

	return res
}

// 动态规划
//1.确定状态
//	最后一步（最优策略中使用的最后一枚硬币Ak）
//	化成子问题（最少的硬币拼出更小的棉质27-Ak）
//2.转移方程
//	f[X]=min{f[X-1]+1, f[X-2]+1, f[X-5]+1}
//3.初始条件和边界情况
//	f[0]=0,若拼不出f[Y], 那么f[Y]=-1
//4.计算顺序
//	f[0],f[1],f[2].....,f[11]
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	//n := len(coins)
	i := 1
	f[0] = 0
	for {
		if i > amount {
			break
		}
		f[i] = math.MaxInt64
		for _, v := range coins {
			if i >= v && f[i-v] != math.MaxInt64 {
				f[i] = int(math.Min(float64(f[i-v]+1), float64(f[i])))
			}
		}
		i++
	}

	if f[amount] == math.MaxInt64 {
		return -1
	}

	return f[amount]
}
