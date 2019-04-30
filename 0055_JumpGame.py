class Solution:
    def canJump(self, nums):
        end = len(nums) - 1
        i = len(nums) - 2
        while i >= 0:
            if i + nums[i] >= end:
                end = i
            i -= 1
        if end == 0:
            return True
        else:
            return False
