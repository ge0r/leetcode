#
# @lc app=leetcode id=15 lang=python
#
# [15] 3Sum
#

# @lc code=start
class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """

        nums.sort()
        i, j, k = 0, 0, 0
        val_to_index = {}

        while (True):
            twosum = nums[i] + nums[k]
            if twosum > 0:
                j = i + 1
            elif twosum < 0:
                j = k-i



# @lc code=end
