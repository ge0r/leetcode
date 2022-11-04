package main

import "sort"

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := []int{}

	i, j := 0, 0

	for {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
		if i == len(nums1) || j == len(nums2) {
			break
		}
	}

	return res
}

// @lc code=end
