class Solution:
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        roman_to_int_dict = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

        sum = 0
        prev = 0
        cur = 0
        for i in s[::-1]:
            cur = roman_to_int_dict[i]

            if prev > cur:
                sum -= cur
            else:
                sum += cur
            prev = cur
        return sum
