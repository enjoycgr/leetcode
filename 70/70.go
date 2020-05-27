//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶

package main

import "fmt"

//1.转移方程， dp[i] = dp[i-1] + dp[i-2]
//2.边界，dp[1] = 1, dp[2] = 2
//3.计算顺序dp[1],dp[2]....dp[3]

func climbStairs(n int) int {
	var dp = make([]int, n+1)

	for i := 1; i < n+1; i++ {
		if i < 3 {
			dp[i] = i
			continue
		}

		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	fmt.Println(climbStairs(3))
}
