class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        #1
        if len(s) != len(t):
            return False

        sC, tC = {}, {}

        for i in range(len(s)):
            sC[s[i]] = 1 + sC.get(s[i], 0)
            tC[t[i]] = 1 + tC.get(t[i], 0)

        for c in sC:
            if sC[c] != tC.get(c, 0):
                return False

        return True

        #2
        # return Counter(s) == Counter(t)

        #3
        # return sorted(s) == sorted(t)
