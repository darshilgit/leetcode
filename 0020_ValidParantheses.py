class Solution:
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        par_dict = {
            ')': '(',
            '}': '{',
            ']': '['
        }
        stack = []
        for paran in s:
            if paran not in par_dict:
                stack.append(paran)
            elif len(stack) > 0 and par_dict[paran] == stack[-1]:
                stack.pop()
            else:
                return False
        if len(stack) > 0:
            return False
        return True
