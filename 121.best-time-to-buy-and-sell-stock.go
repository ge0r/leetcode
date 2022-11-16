package leetcode

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	min := 10000
	profit := 0

	for _, price := range prices {
		if (price - min) > profit {
			profit = price - min
		}
		if price < min {
			min = price
		}
	}

	return profit
}

// @lc code=end
