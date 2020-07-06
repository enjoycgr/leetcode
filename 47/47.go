//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//
//示例:
//
//输入: [1,1,2]
//输出:
//[
//	[1,1,2],
//	[1,2,1],
//	[2,1,1]
//]

package main

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	used := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		used[nums[i]]++
	}
	help(nums, []int{}, []int{})
}

func help(nums, output, used []int) {
	if len(output) == len(nums) {

	}

	for i := 0; i < len(nums); i++ {
		if used[nums[i]] == 0 {
			output = append(output, )
		}
	}
}

func main() {
	
}
