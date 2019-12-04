//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//	[-1, 0, 1],
//	[-1, -1, 2]
//]

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var left, right int
	var result [][]int
	var sum int
	length := len(nums)

	if length < 3 {
		return result
	}
	for k, v := range nums {
		left = k + 1
		right = len(nums) - 1
		// 避免重复
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		// 剩下2个退出
		if k > right-2 {
			break
		}
		for {
			if left >= right {
				break
			}
			sum = nums[left] + nums[right] + v
			if sum == 0 {
				result = append(result, []int{nums[left], nums[right], v})
				// 去重
				for {
					if left < right && nums[left] == nums[left+1] {
						left = left + 1
					} else {
						break
					}
				}
				// 去重
				for {
					if left < right && nums[right] == nums[right-1] {
						right = right - 1
					} else {
						break
					}
				}
				left = left + 1
				right = right - 1
			}

			if sum > 0 {
				right = right - 1
			}
			if sum < 0 {
				left = left + 1
			}
		}
	}
	return result
}
