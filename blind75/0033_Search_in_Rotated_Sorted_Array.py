class Solution:

    def findPivot(self, nums):
        l, r = 0, len(nums) - 1

        while l < r:
            mid = (l+r) // 2
            if nums[mid] > nums[r]:
                l = mid + 1
            else:
                r = mid

        return r

    def binarySearch(self, l, r, nums, target):
        idx = -1
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] == target:
                return mid
            if nums[mid] > target:
                r = mid - 1
            else:
                l = mid + 1

        return idx

    def search(self, nums: List[int], target: int) -> int:
        pivot = self.findPivot(nums)

        if nums[pivot] == target:
            return pivot

        idx = self.binarySearch(0, pivot-1, nums, target)

        if idx != -1:
            return idx

        return self.binarySearch(pivot, len(nums) - 1, nums, target)