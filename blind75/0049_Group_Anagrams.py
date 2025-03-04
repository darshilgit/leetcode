class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = defaultdict(list)

        for string in strs:
            alphaCount = [0] * 26

            for char in string:
                idx = ord(char) - ord('a')
                alphaCount[idx] += 1

            res[tuple(alphaCount)].append(string)

        return list(res.values())