class Solution:
    def permute(self, nums):
        answer = []

        self.permuteHelper(0, nums, answer)
        return answer

    def permuteHelper(self, index, nums, answer):
        if index == len(nums) - 1:
            answer.append(nums.copy())

        for j in range(index, len(nums)):
            nums[index], nums[j] = nums[j], nums[index]
            self.permuteHelper(index + 1, nums, answer)
            nums[j], nums[index] = nums[index], nums[j]
