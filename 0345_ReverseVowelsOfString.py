class Solution:
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        vow_dict = {
            "a": 1,
            "A": 1,
            "e": 2,
            "E": 2,
            "i": 3,
            "I": 3,
            "o": 4,
            "O": 4,
            "u": 5,
            "U": 5
        }
        str_len = len(s)
        vow_list = []
        i = 0
        for letter in s:
            if letter in vow_dict.keys():
                vow_list.append(letter)
                i += 1
        vow_list = vow_list[::-1]
        x = 0
        answer = ""
        for ele in range(str_len):
            if s[ele] not in vow_dict :
                answer += s[ele]
            else:
                answer += vow_list[x]
                x += 1
        return answer


a = Solution()
print(a.reverseVowels("leetcode"))