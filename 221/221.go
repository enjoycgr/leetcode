//在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
//
//示例:
//
//输入:
//
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//
//输出: 4

package main

import (
	"fmt"
	"math"
)

// 1. 确定状态：dp[i][j]确定以matrix[i][j]为右下角元素的正方向边长
// 2. 转移方程：当matrix[i][j] == 1时，dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
//   		  当matrix[i][j] == 0时，dp[i][j] = 0;
// 3. 边界情况：当i = 0 || j = 0, dp[i][j] = matrix[i][j]
// 4. 计算顺序：dp[0][0],dp[0][1]....dp[n][n]

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for k, _ := range dp {
		dp[k] = make([]int, len(matrix[k]))
	}
	max := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j]) - 48
			} else if int(matrix[i][j]) == 49 { // matrix[i][j]为1
				min := int(math.Min(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])), float64(dp[i-1][j-1])))
				dp[i][j] = 1 + int(min)
			} else {
				dp[i][j] = 0
			}

			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max * max
}

func main() {
	res := maximalSquare([][]byte{
		[]byte{'1'},
	})
	fmt.Println(res)
}
