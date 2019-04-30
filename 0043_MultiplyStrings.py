class Solution:
    def multiply(self, num1, num2):
        t1 = []
        l1 = len(num1)
        l2 = len(num2)
        if l1 == 1 and l2 == 1:
            return str(int(num1) * int(num2))
        pad_count = 0
        # if l2 > l1:
        #     temp = num2
        #     num2 = num1
        #     num1 = temp
        #     l1 = len(num1)
        #     l2 = len(num2)

        for i in range(l2 - 1, -1, -1):
            c = 0
            r = ""
            for j in range(l1 - 1, -1, -1):
                d1 = int(num2[i])
                d2 = int(num1[j])
                p = d2 * d1 + c
                digit = int(p % 10)
                c = int(p // 10)
                r = str(digit) + r

            if c > 0:
                r = str(c) + str(r)

            if pad_count > 0:
                pad = "0" * pad_count
                r = r + pad
            t1.append(int(r))
            pad_count = pad_count + 1

        result = str(sum(t1))
        return result