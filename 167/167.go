//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//
//函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//
//说明:
//
//返回的下标值（index1 和 index2）不是从零开始的。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
//示例:
//
//输入: numbers = [2, 7, 11, 15], target = 9
//输出: [1,2]
//解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

package main

func twoSum(numbers []int, target int) []int {
	for k, _ := range numbers {
		if numbers[k] > target {
			continue
		}
		res := help(numbers, k, k + 1, target)
		if res > 0 {
			return []int{k + 1, res + 1}
		}
	}

	return nil
}

func help(number []int, start, end, target int) int {
	if number[start] + number[end] == target {
		return end
	}
	if end == len(number) - 1 || number[end] > target {
		return -1
	}

	return help(number, start, end + 1, target)
}

func main() {
	
}
