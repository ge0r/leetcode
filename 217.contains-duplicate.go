package main

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)
	for _, num := range nums {
		if _, ok := visited[num]; ok {
			return true
		} else {
			visited[num] = true
		}
	}
	return false
}

// @lc code=end
