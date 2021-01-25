class Solution:
    def regionsBySlashes(self, grid: List[str]) -> int:
        n = len(grid)                

        f = [i for i in range(n*n*4)]
        def getF(x):
            f[x] = x if x == f[x] else getF(f[x])
            return f[x]
        def connect(x, y):
            f[getF(x)] = getF(y)
        def index(i, j, k):
            return i * n * 4 + j * 4 + k

        # |——————————|
        # |  \  0  / |
        # | 3   X  1 |
        # |  /  2 \  |
        # |——————————|


        for i in range(n):
            for j in range(n):
                c = grid[i][j]
                if c == ' ':
                    # 空白
                    connect(index(i, j, 0), index(i, j, 1))
                    connect(index(i, j, 0), index(i, j, 2))
                    connect(index(i, j, 0), index(i, j, 3))
                elif c == '/':
                    # /
                    connect(index(i, j, 0), index(i, j, 3))
                    connect(index(i, j, 1), index(i, j, 2))
                else:
                    # \
                    connect(index(i, j, 0), index(i, j, 1))
                    connect(index(i, j, 2), index(i, j, 3))

                if i != 0:
                    connect(index(i, j, 0), index(i-1, j, 2))
                if j != 0:
                    connect(index(i, j, 3), index(i, j-1, 1))

        res = 0
        for i in range(n*n*4):
            if getF(i) == i:
                res += 1
        # print(f)

        return res



