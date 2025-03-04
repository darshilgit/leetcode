class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        maxProd = max(nums) # cant be zero since if only one element and thats -ve, it will be that value

        currMin = 1
        currMax = 1

        for n in nums:
            if n == 0:
                currMin, currMax = 1, 1
                continue

            tmpMax = n * currMax # to avoid using overridden value from below line
            currMax = max(n* currMax, n * currMin, n) # for just n -> [-1, 8]
            currMin = min(tmpMax, n * currMin, n) # for just n -> [-1, -8]

            maxProd = max(currMax, maxProd)
        
        return maxProd

        # Time: O(n)
        # Space: O(1)
