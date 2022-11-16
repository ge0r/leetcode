package leetcode

/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	size := len(nums)
	left, right := 0, size-1
	res := make([]int, size)

	for i := size - 1; i >= 0; i-- {
		lres := nums[left] * nums[left]
		rres := nums[right] * nums[right]
		if lres > rres {
			res[i] = lres
			left++
		} else {
			res[i] = rres
			right--
		}
	}
	return res
}

// @lc code=end
