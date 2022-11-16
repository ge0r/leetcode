#
# @lc app=leetcode id=977 lang=python3
#
# [977] Squares of a Sorted Array
#

# @lc code=start
class Solution:
    def sortedSquares(self, nums: list[int]) -> list[int]:
        n = len(nums)
        res = [0] * n
        i, j, k = 0, n-1, n-1

        while k >= 0:
            if abs(nums[i]) > abs(nums[j]):
                res[k] = nums[i] * nums[i]
                i += 1
            else:
                res[k] = nums[j] * nums[j]
                j -= 1

            k-=1
        return res

# @lc code=end
