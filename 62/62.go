//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//问总共有多少条不同的路径？
//输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右

// 动态规划
//1.确定状态
//	最后一步f[i-1][j],f[i][j-1]
//	化成子问题（有多少条不同的路径到达f[i-1][j]和f[i][j-1]）
//2.转移方程
//	f[i][j]=f[i-1][j]+f[i][j-1]
//3.初始条件和边界情况
//	初始条件:f[0][0]=1
//	边界情况：i=0或j=0,前一步只能从一个方向走来,这时候f[i][j]=1
//4.计算顺序
//	计算第0行：f[0][0],f[0][1],...,f[0][m]
//	计算第1行：f[1][0],f[1][1],...,f[1][m]
//  ...
//	计算第n行: f[n][0],f[n][1],...,f[n][m]

package main

func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				f[i][j] = 1
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}

	return f[m-1][n-1]
}