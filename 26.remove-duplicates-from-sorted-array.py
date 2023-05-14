#
# @lc app=leetcode id=26 lang=python3
#
# [26] Remove Duplicates from Sorted Array
#

# @lc code=start
class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        i, j = 1, 1
        while (j < len(nums)):
            if nums[j] != nums[j-1]:
                nums[i] = nums[j]
                i+=1
            j+=1
        return i

# @lc code=end
