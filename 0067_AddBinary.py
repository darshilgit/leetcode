class Solution:
    def addBinary(self, a, b):
        res = ''
        i = len(a) - 1
        j = len(b) - 1

        carry = 0

        while i >= 0 or j >= 0:
            _a = int(a[i]) if i >= 0 else 0
            _b = int(b[j]) if j >= 0 else 0
            if _a + _b + carry == 1:
                _sum = 1
                carry = 0
            elif _a + _b + carry == 2:
                _sum = 0
                carry = 1
            elif _a + _b + carry > 2:
                _sum = 1
                carry = 1
            else:
                _sum = 0
                carry = 0
            res = str(_sum) + res
            i -= 1
            j -= 1
        if carry == 1:
            res = str(carry) + res
        return res