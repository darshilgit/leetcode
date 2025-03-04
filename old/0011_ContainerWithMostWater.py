class Solution:
    def maxArea(self, height):
        left = 0
        right = len(height) - 1

        curr_max_area = 0
        while left < right:
            if min(height[left], height[right]) * (right - left) > curr_max_area:
                curr_max_area = min(height[left], height[right]) * (right - left)

            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return curr_max_area