package leetcode

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j := 0, 0
	res := []int{}

	for i < m || j < n {
		if i < m {
			if j < n {
				if nums1[i] < nums2[j] {
					res = append(res, nums1[i])
					i++
				} else {
					res = append(res, nums2[j])
					j++
				}
			} else {
				res = append(res, nums1[i])
				i++
			}
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	copy(nums1, res)
}

// @lc code=end
