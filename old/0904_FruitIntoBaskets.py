class Solution:
    def totalFruit(self, tree):
        """
        :type tree: List[int]
        :rtype: int
        """
        fruit_dict = {}
        left = 0
        right = 0
        maximum = 0
        while right < len(tree):
            if len(fruit_dict.keys()) <= 2:
                fruit_dict[tree[right]] = right
                right += 1
            if len(fruit_dict.keys()) > 2:
                minimum = len(tree) - 1
                for ele in fruit_dict.values():
                    minimum = min(minimum, ele)
                left = minimum + 1
                del fruit_dict[tree[minimum]]
            maximum = max(maximum, right - left)

        return maximum

a = Solution()
print(a.totalFruit([1,2,3,2,2]))
