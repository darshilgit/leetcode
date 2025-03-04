class Solution:
    def majorityElement(self, nums):
        nums.sort()
        if len(nums) == 1:
            return nums[0]
        return nums[int(len(nums) / 2)]