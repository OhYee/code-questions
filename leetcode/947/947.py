class Solution:
    def removeStones(self, stones: List[List[int]]) -> int:
        n = len(stones)
        f = [i for i in range(n)]
        def getF(x):
            f[x] = x if f[x] == x else getF(f[x])
            return f[x]
        def connect(i, j):
            f[getF(i)] = getF(j)
        
        dx = {}
        dy = {}
        for i, stone in enumerate(stones):
            x, y = stone

            sx = dx.get(x, None)
            if sx == None:
                dx[x] = i
            else:
                connect(i, sx)
            
            sy = dy.get(y, None)
            if sy == None:
                dy[y] = i
            else:
                connect(i, sy)
        
        count = 0
        for i in range(n):
            if getF(i) == i:
                count += 1

        return n - count
            


