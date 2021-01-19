import queue

class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        n = len(points)
        
        q = queue.PriorityQueue()
        for i in range(n):
            for j in range(i + 1, n):
                a = points[i]
                b = points[j]
                d = abs(a[0] - b[0]) + abs(a[1] - b[1])
                q.put([d, i, j])
        
        f = [i for i in range(n)]
        c = [1 for i in range(n)]
        def getF(x):
            f[x] = x if x == f[x] else getF(f[x])
            return f[x]
        def connect(a, b):
            root_a = getF(a)
            root_b = getF(b)
            c[root_a] = c[root_a] + c[root_b]
            c[root_b] = c[root_a]
            f[getF(a)] = getF(b)
            return c[root_a]

        res = 0
        while not q.empty():
            d, i, j = q.get()
            root_i = getF(i)
            root_j = getF(j)
            
            if root_i == root_j:
                # 两个点已联通
                continue
            res += d
            if connect(i, j) >= n:
                break
            
        return res




