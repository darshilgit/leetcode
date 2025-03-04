import re
class Solution:
    def isPalindrome(self, s):
        s = re.sub('\W', '', s).lower()

        left = 0
        right = len(s) - 1

        while (left < right):
            if s[left] == s[right]:
                left += 1
                right -= 1
            else:
                return False
        return True