class Solution:
    def maxNumEdgesToRemove(self, n: int, edges: List[List[int]]) -> int:
        edges = sorted(edges, reverse=True)

        aliceF = [i for i in range(n)]
        bobF = [i for i in range(n)]
        def getF(f, x):
            f[x] = x if x == f[x] else getF(f, f[x])
            return f[x]
        def connect(f, x, y):
            f[getF(f, x)] = getF(f, y)
        

        res = 0
        for t, x, y in edges:
            x -= 1
            y -= 1
            if t == 1:
                rootX = getF(aliceF, x)
                rootY = getF(aliceF, y)
                if rootX != rootY:
                    res += 1
                    connect(aliceF, x, y)
            elif t == 2:
                rootX = getF(bobF, x)
                rootY = getF(bobF, y)
                if rootX != rootY:
                    res += 1
                    connect(bobF, x, y)
            else:
                rootX = getF(aliceF, x)
                rootY = getF(aliceF, y)
                if rootX != rootY:
                    res += 1
                    connect(aliceF, x, y)
                    connect(bobF, x, y)
        ac = 0
        bc = 0
        for i in range(n):
            if aliceF[i] == i:
                ac += 1
            if bobF[i] == i:
                bc += 1
        return len(edges) - res if ac == 1 and bc == 1 else -1
