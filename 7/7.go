//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//
//输入: 123
//输出: 321
// 示例 2:
//
//输入: -123
//输出: -321
//示例 3:
//
//输入: 120
//输出: 21

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	i := true
	if x < 0 {
		i = false
		x = 0 - x
	}
	s := 0
	for x > 0 {
		s = s*10 + x%10
		x = x / 10
	}

	if i == false {
		s = -s
		if s < math.MinInt32 {
			return 0
		}
	}

	if s > math.MaxInt32 {
		return 0
	}

	return s
}
