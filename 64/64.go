//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
//示例:
//
//输入:
//[
//[1,3,1],
//[1,5,1],
//[4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。

package main

import "fmt"

// 1.转移方程 dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
// 2.边界， i = 0时， dp[i][j] = dp[i][j-1] + grid[i][j]
//         j = 0时,  dp[i][j] = dp[i-1][j] + grid[i][j]
// 3.初始， dp[0][0] = grid[i][j]
// 4.计算顺序, dp[0][0],dp[0][1]....dp[m][n]

func minPathSum(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
				continue
			}

			if j == 0 && i != 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
				continue
			}

			if grid[i][j-1] < grid[i-1][j] {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			}
		}
	}

	return grid[m-1][n-1]
}

func main() {
	data := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	fmt.Println(minPathSum(data))
}
