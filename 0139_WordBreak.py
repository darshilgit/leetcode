class Solution:
    def wordBreak(self, s, wordDict):
        dp = [False for _ in range(len(s)+1)]
        word_set = set()
        for word in wordDict:
            word_set.add(word)
        dp[0] = True
        for i in range(0,len(s)+1):
            for j in range(0,i):
                if dp[j] and s[j:i] in word_set:
                    dp[i] = True
                    break
        return dp[-1]