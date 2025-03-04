class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if not grid:
            return 0

        rows, cols = len(grid), len(grid[0])
        visit = set()
        numIslands = 0

        def dfs(r, c):
            st = []
            st.append((r,c))
            visit.add((r,c))

            while st:
                row, col = st.pop()
                # neighbors
                directions = [[0, 1], [0, -1], [1, 0], [-1, 0]]

                for dr, dc in directions:
                    rdr, cdc = row+dr, col+dc
                    if rdr in range(rows) and cdc in range(cols) and grid[rdr][cdc] == '1' and (rdr, cdc) not in visit:
                        st.append((rdr, cdc))
                        visit.add((rdr, cdc))

        def bfs(r, c):
            q = deque()
            q.append((r,c))
            visit.add((r,c))

            while q:
                row, col = q.popleft()
                directions = [[0, 1], [0, -1], [1, 0], [-1, 0]]

                for dr, dc in directions:
                    rdr, cdc = row+dr, col+dc
                    if rdr in range(rows) and cdc in range(cols) and grid[rdr][cdc] == '1' and (rdr, cdc) not in visit:
                        q.append((rdr, cdc))
                        visit.add((rdr, cdc))

        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == '1' and (r,c) not in visit:
                    dfs(r,c)
                    numIslands += 1

        return numIslands
