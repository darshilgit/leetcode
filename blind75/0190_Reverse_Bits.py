class Solution:
    def reverseBits(self, n: int) -> int:
        rev_n = 0
        for i in range(32):
            bit = n & 1
            rev_n = rev_n << 1
            rev_n = rev_n | bit
            n = n >> 1
        return rev_n
