class Solution:
    def maximalSquare(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if len(matrix) == 0:
            return 0
        elif len(matrix[0]) == 0:
            return 0
        cols = len(matrix[0])
        rows = len(matrix)
        cache = [[0 for i in range(cols)] for j in range(rows)]
        maximum = 0
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                if j == 0:
                    cache[i][j] = int(matrix[i][j])
                elif int(matrix[i][j]) == 1:
                    cache[i][j] = min(min(int(cache[i][j - 1]), int(cache[i - 1][j])), int(cache[i - 1][j - 1])) + 1
                maximum = max(maximum, cache[i][j])
        return maximum*maximum


a = Solution()
print(a.maximalSquare([["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]))