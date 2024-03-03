package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	mid := (len(xStr) / 2)

	for i := 0; i < mid; i++ {
		if string(xStr[i]) != string(xStr[len(xStr)-i-1]) {
			return false
		}
	}
	return true
}

// @lc code=end
