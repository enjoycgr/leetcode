//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2


package main

import "fmt"

// 分治法
// 如果x是数组nums众数的话,x出现次数a>n/2，x肯定是[0,n/2]或[n/2+1,n]的众数
// 反证法：定义left为[0,n/2]的长度，right为[n/2+1,n]的长度。
//        如果x不是[0,n/2]和[n/2+1,n]的众数， a <= l/2 + r/2， 得a <= n/2，那么a必然不是nums的众数

func main() {
	fmt.Println(majorityElement([]int{2,2,1,1,1,1,2,2}))
}

func majorityElement(nums []int) int {
	return majorityElementRec(nums, 0, len(nums) - 1)
}

func majorityElementRec(nums []int, left, right int) int {
	// 只有一个数的时候必定是众数
	if left == right {
		return nums[left]
	}

	mid := int((right - left) / 2 + left)
	// 左右两部分的众数
	lElement := majorityElementRec(nums, left, mid)
	rElement := majorityElementRec(nums, mid+1, right)

	// 左右两部分众数一样时返回
	if lElement == rElement {
		return lElement
	}

	// 计算左右两部分的众数在当前数组中的出现次数并得出当前数组的众数
	leftCount := countInRange(nums, lElement, left, right)
	rightCount := countInRange(nums, rElement, left, right)

	if leftCount > rightCount {
		return lElement
	}
	return rElement
}

func countInRange(nums []int, num, left, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if nums[i] == num {
			count++
		}
	}

	return count
}