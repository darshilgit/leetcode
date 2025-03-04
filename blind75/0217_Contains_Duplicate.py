class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        cdset = set()

        for n in nums:
            if n in cdset:
                return True
            else:
                cdset.add(n)
        
        return False
