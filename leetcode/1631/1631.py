import queue

class Solution:
    def minimumEffortPath(self, heights: List[List[int]]) -> int:
        n = len(heights)
        if n == 0:
            return 0
        m = len(heights[0])

        vis = [[ False for j in range(m)] for i in range(n)]

        q = queue.PriorityQueue()
        q.put((0, 0, 0))

        while not q.empty():
            d, x, y = q.get()
            # print(d,x,y)
            if x == n-1 and y == m-1:
                return d
            if vis[x][y] == False:
                vis[x][y] = True
            else:
                continue
            if x > 0:
                xx = x-1
                yy = y
                dd = abs(heights[xx][yy] - heights[x][y])
                q.put((max(d,dd), xx, yy))
            if y > 0:
                xx = x
                yy = y-1
                dd = abs(heights[xx][yy] - heights[x][y])
                q.put((max(d,dd), xx, yy))
            if x+1 < n:
                xx = x+1
                yy = y
                dd = abs(heights[xx][yy] - heights[x][y])
                q.put((max(d,dd), xx, yy))
            if y+1 < m:
                xx = x
                yy = y+1
                dd = abs(heights[xx][yy] - heights[x][y])
                q.put((max(d,dd), xx, yy))
        return 0
