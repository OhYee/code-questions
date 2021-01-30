import queue

class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        n = len(grid)

        vis = [[False for j in range(n)] for i in range(n)]
        q = queue.PriorityQueue()

        q.put((grid[0][0], 0, 0))
        while not q.empty():
            h, x, y = q.get()
            if x == n-1 and y == n-1:
                return h
            if vis[x][y] == True:
                continue
            vis[x][y] = True

            for dx, dy in [(-1, 0), (0, -1), (1, 0), (0, 1)]:
                xx = x + dx
                yy = y + dy
                if 0 <= xx and xx < n and 0 <= yy and yy < n:
                    hh = max(grid[xx][yy], h)
                    q.put((hh, xx, yy))
        return -1

