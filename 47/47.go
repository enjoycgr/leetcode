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
		// 选择路径
		if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
			continue
		}
		if !used[i] {
			current := nums[i]
			// 添加
			used[i] = true
			output = append(output, nums[i])
			help(used, nums, output)
			// 回溯
			output = output[:len(output) - 1]
			used[i] = false

			for j := i + 1; j < len(nums); j++ {
				if nums[j] == current {
					i++
				} else {
					break
				}
			}
		}
	}
}


