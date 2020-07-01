//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//	[1,2,3],
//	[1,3,2],
//	[2,1,3],
//	[2,3,1],
//	[3,1,2],
//	[3,2,1]
//]

package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{} // 最终结果集
	used := make([]bool, len(nums)) // 标记数字是否被使用过
	help(used, nums, []int{})
	return res
}

func help(used []bool, nums, output []int) {
	if len(nums) == len(output) {
		tmp := make([]int, len(nums))
		copy(tmp, output)
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			output = append(output, nums[i])
			help(used, nums, output)
			output = output[0:len(output) - 1]
			used[i] = false
		}
	}
}

func main() {
	a := []int{1}
	test(a)
	fmt.Println(a)
}

func test(a []int) {
	a = []int{2}
}