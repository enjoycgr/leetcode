package main

func merge(nums []int, left, right int) []int {
	if left == right {
		return nums[left : left+1]
	}

	mid := (right-left)/2 + left

	lsort := merge(nums, left, mid)
	rsort := merge(nums, mid+1, right)

	res := msort(lsort, rsort)

	return res
}

func msort(left, right []int) []int {
	var l, r = 0, 0
	var s []int
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			s = append(s, left[l])
			l++
			continue
		}
		if left[l] <= right[r] {
			s = append(s, right[r])
			r++
			continue
		}
	}

	if l == len(left) {
		for i := r; i < len(right); i++ {
			s = append(s, right[i])
		}
	} else {
		for i := l; i < len(left); i++ {
			s = append(s, left[i])
		}
	}

	return s
}
