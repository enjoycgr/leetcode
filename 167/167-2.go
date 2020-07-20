package main

// 双指针
func twoSum2(numbers []int, target int) []int {
	return help2(numbers, 0, len(numbers) - 1, target)
}

func help2(numbers []int, l, r, target int) []int {
	if numbers[l] + numbers[r] == target {
		return []int{l + 1, r + 1}
	}

	if l == r {
		return nil
	}

	if numbers[l] + numbers[r] > target {
		return help2(numbers, l, r - 1, target)
	} else {
		return help2(numbers, l + 1, r, target)
	}
}

