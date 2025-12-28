class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        # helper function to found nearby land
        def get_neighbours(i, j):
            res = []
            di = [0, 1, -1, 0]
            dj = [-1, 0, 0, 1]
            for idx in range(4):
                new_i = i + di[idx]
                new_j = j + dj[idx]
                if 0 <= new_i < len(grid[0]) and 0 <= new_j < len(grid):
                    if grid[new_j][new_i] == "1":
                        res.append((new_i, new_j))
            return res

        # turn every island to water to dont count visited ones
        def dfs(i, j):
            if grid[j][i] == "0":
                return
            grid[j][i] = "0"
            neighbours = get_neighbours(i, j)
            for x, y in neighbours:
                if grid[y][x] == "0":
                    continue
                dfs(x, y)
            

        # go through every cell to find land ad count them
        count = 0
        for j in range(len(grid)):
            for i in range(len(grid[0])):
                if grid[j][i] == "1":
                    dfs(i, j)
                    count += 1
        return count