package main

import "fmt"

func main() {
	fmt.Println(quick([]int{2, 5, 18, 99, 1, 23, 1, 3}, 0, 7))
}

func quick(nums []int, left, right int) []int {
	mid := qsort(nums, left, right)

	if left < right {
		quick(nums, left, mid-1)
		quick(nums, mid+1, right)
	}

	return nums
}

func qsort(nums []int, left, right int) int {
	if left > right {
		return left
	}
	key := nums[left]
	for left < right {
		if left < right && nums[right] >= key {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
		if left < right && nums[left] <= key {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}

	return left
}
