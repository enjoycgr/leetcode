//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//你可以假设 nums1 和 nums2 不会同时为空。
//
//示例 1:
//nums1 = [1, 3]
//nums2 = [2]
//则中位数是 2.0

package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	count := len(nums1) + len(nums2)
	var n1, n2, num int

	if count%2 > 0 {
		for {
			if n1 < len(nums1) && n2 < len(nums2) && nums1[n1] < nums2[n2] {
				if n1+n2 == count/2 {
					return float64(nums1[n1])
				}
				n1++
			}
			if n1 < len(nums1) && n2 < len(nums2) && nums2[n2] <= nums1[n1] {
				if n1+n2 == count/2 {
					return float64(nums2[n2])
				}
				n2++
			}

		}
	} else {
		for {
			if n1 < len(nums1) && n2 < len(nums2) && nums1[n1] < nums2[n2] {
				if n1+n2 == count/2 {
					return (float64(nums1[n1]) + float64(num)) / 2
				}
				num = nums1[n1]
				n1++
			}
			if n1 < len(nums1) && n2 < len(nums2) && nums2[n2] <= nums1[n1] {
				if n1+n2 == count/2 {
					return (float64(nums2[n2]) + float64(num)) / 2
				}
				num = nums2[n2]
				n2++
			}

		}
	}
}
