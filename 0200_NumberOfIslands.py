class Solution:
    def numIslands(self, grid):
        if grid == None or len(grid) == 0:
            return 0
        num_islands = 0
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                num_islands += self.dfs(grid, i, j)

        return num_islands

    def dfs(self, grid, i, j):
        if i < 0 or i >= len(grid) or j >= len(grid[i]) or j < 0 or grid[i][j] == '0':
            return 0
        grid[i][j] = '0'

        self.dfs(grid, i - 1, j)
        self.dfs(grid, i + 1, j)
        self.dfs(grid, i, j - 1)
        self.dfs(grid, i, j + 1)

        return 1
