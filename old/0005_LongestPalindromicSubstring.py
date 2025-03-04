class Solution:
    def longestPalindrome(self, s):
        max_length = 1

        start = 0

        for i in range(1, len(s)):
            # Checking for even length palindrome
            l = i - 1
            h = i
            while l >= 0 and h < len(s) and s[l] == s[h]:
                if h - l + 1 > max_length:
                    start = l
                    max_length = h - l + 1
                l -= 1
                h += 1
            # Checking for odd length palindromes
            l = i - 1
            h = i + 1
            while l >= 0 and h < len(s) and s[l] == s[h]:
                if h - l + 1 > max_length:
                    start = l
                    max_length = h - l + 1
                l -= 1
                h += 1
        return s[start:start + max_length]
