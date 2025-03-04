import heapq

class Solution(object):
    def findKthLargest(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        # Approach 1
        # nums.sort()
        # return nums[-k]

        # Approach 2
        return heapq.nlargest(k, nums)[-1]

