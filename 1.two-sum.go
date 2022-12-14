package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	visited := make(map[int]int)

	for i, num := range nums {
		if j, ok := visited[target-num]; ok {
			return []int{j, i}
		} else {
			visited[num] = i
		}
	}
	return []int{}
}

// @lc code=end
