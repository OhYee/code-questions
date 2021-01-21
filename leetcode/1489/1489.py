import queue

class Solution:
    def findCriticalAndPseudoCriticalEdges(self, n: int, edges: List[List[int]]) -> List[List[int]]:
        self.n = n
        self.m = len(edges)
        self.f = [i for i in range(n)]
        self.c = [1 for i in range(n)]
        
        # 先生成一个最小生成树
        mstWeight, mstEdges = self.getTree(n, edges)

        # 移除所有权值超过 MST 最大权值的边
        # 这些边必然不会出现在 MST 中
        maxWeight = 0
        for e in mstEdges:
            maxWeight = max(maxWeight, edges[e][2])

        edgeType = [0 for i in range(self.m)] # 边类型 0 未知; 1 关键; 2 伪关键; 3 无用边
        for i in range(self.m):
            if edges[i][2] > maxWeight:
                edgeType[i] = 3

        # 关键边必然是这个最小生成树的一条边
        # 检查是否为关键边: 移除这条边是否还能生成 MST
        for e in mstEdges:
            w, r = self.getTree(n, edges, ignoreEdge=e)
            if w == mstWeight:
                # 伪关键边
                edgeType[e] = 2
            else:
                # 关键边
                edgeType[e] = 1

        # 判断所有剩余的边是否为伪关键边
        # 这些边必然不是关键边，只可能是伪关键边或无用边
        # 只需要判断添加这些边是否能构成 MST 即可
        for e in range(self.m):
            if edgeType[e] == 0:
                w, r = self.getTree(n, edges, mustEdge=e)
                if w == mstWeight:
                    # 可以生成 MST
                    edgeType[e] = 2
                else:
                    # 无法生成 MST
                    edgeType[e] = 3

        # 生成结果数组
        keyEdges = []
        fakeKeyEdges = []
        for e in range(self.m):
            if edgeType[e] == 1:
                keyEdges.append(e)
            elif edgeType[e] == 2:
                fakeKeyEdges.append(e)

        return keyEdges, fakeKeyEdges

    def getF(self, x):
        self.f[x] = x if x == self.f[x] else self.getF(self.f[x])
        return self.f[x]

    def connect(self, x, y):
        rootX = self.getF(x)
        rootY = self.getF(y)
        self.c[rootX] = self.c[rootX] + self.c[rootY]
        self.c[rootY] = self.c[rootX]
        self.f[rootX] = self.getF(rootY)
        return self.c[rootX]

    def getTree(self, n, edges, ignoreEdge=-1, mustEdge=-1):
        for i in range(self.n):
            self.f[i] = i
            self.c[i] = 1

        q = queue.PriorityQueue()
        for i, e in enumerate(edges):
            if i != mustEdge:
                q.put((e[2], i))            

        dis = 0
        res = []
        if mustEdge != -1:
            x, y, d = edges[mustEdge]
            dis += d
            self.connect(x, y)
            res.append(mustEdge)
        while not q.empty():
            d, t = q.get()
            if t == ignoreEdge:
                continue
            x, y, d = edges[t]
            rootX = self.getF(x)
            rootY = self.getF(y)
            if rootX != rootY:
                res.append(t)
                dis += d
                if self.connect(rootX, rootY) == n:
                    break
        return dis, res
