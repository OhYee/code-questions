class Solution:
    def hitBricks(self, grid: List[List[int]], hits: List[List[int]]) -> List[int]:
        if len(grid) == 0:
            return [0] * len(hits)

        # import time
        # start = time.time()
        
        n = len(grid)
        m = len(grid[0])

        ngrid = [[grid[x][y] for y in range(m)] for x in range(n)]
        for h in hits:
            x, y = h
            ngrid[x][y] = 0

        # print(time.time() - start)

        # for x in range(n):
        #     for y in range(m):
        #         print(ngrid[x][y], end=" ")
        #     print()

        f = [i for i in range(m * n + 1)]
        c = [1 for i in range(m * n + 1)]
        def getIdx(x, y):
            return x * m + y + 1   
        def getF(t):
            f[t] = f[t] if t == f[t] else getF(f[t])
            return f[t]
        def getC(t):
            # c[t] = c[t] if t == f[t] else getC(f[t])
            return c[getF(t)]
        def connect(t1, t2):
            root1 = getF(t1)
            root2 = getF(t2)
            if root1 != root2:
                c[root2] = c[root1] + c[root2]
                c[root1] = c[root2]
                f[root1] = root2

        for x in range(n):
            for y in range(m):
                if ngrid[x][y] == 1:
                    t = getIdx(x, y)
                    for xx, yy in [(x+1,y), (x,y+1)]:
                        if 0 <= xx and xx < n and 0 <= yy and yy < m and ngrid[xx][yy] == 1:
                            tt = getIdx(xx, yy)
                            connect(t, tt)
                            # print("connect", t,x,y, tt,xx,yy)
                    if x == 0:
                        connect(0, t)

        # print(time.time() - start)

        # for x in range(n):
        #     for y in range(m):
        #         print(getF(getIdx(x, y)), end=" ")
        #     print()

        # for x in range(n):
        #     for y in range(m):
        #         print(getC(getIdx(x, y)), end=" ")
        #     print()


        res = [0] * len(hits)
        for i, hit in enumerate(reversed(hits)):
            x, y = hit
            if grid[x][y] == 1:
                ngrid[x][y] = 1
                beforeNum = getC(0)
                # print("insert",x,y)

                t = getIdx(x, y)
                for xx, yy in [(x-1,y),(x+1,y),(x,y-1),(x,y+1)]:
                    if 0 <= xx and xx < n and 0 <= yy and yy < m and ngrid[xx][yy] == 1:
                        tt = getIdx(xx, yy)
                        connect(t, tt)
                        # print("connect2", t,x,y, tt,xx,yy)
                if x == 0:
                    connect(0, t)

                nowNum = getC(0)
                # print("before", beforeNum, "now", nowNum)
                res[i] = max(nowNum - beforeNum - 1, 0)

        # print(time.time() - start)

        res.reverse()
        return res



            
    
