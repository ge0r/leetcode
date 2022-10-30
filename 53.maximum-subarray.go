package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	maxSum := -10000
	sum := 0

	for _, num := range nums {
		sum += num
		if num > sum {
			sum = num
		}
		if sum >= maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

// @lc code=end
