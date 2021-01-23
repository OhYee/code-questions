class Solution:
    def makeConnected(self, n: int, connections: List[List[int]]) -> int:
        f = [i for i in range(n)]
        def getF(x):
            f[x] = x if x == f[x] else getF(f[x])
            return f[x]
        def connect(a, b):
            f[getF(a)] = getF(b)

        count = 0
        for c in connections:
            a, b = c
            rootA = getF(a)
            rootB = getF(b)
            if rootA == rootB:
                count += 1
            else:
                connect(rootA, rootB)
        
        group = 0
        for i in range(n):
            if f[i] == i:
                group += 1
        
        if group - 1 <= count:
            return min(group - 1, count)
        return -1
