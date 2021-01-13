class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        n = len(edges)
        
        self.e = [[] for i in range(n+1)]
        self.v = [False for i in range(n+1)]
        self.res = []

        for edge in edges:
            a, b = edge
            self.e[a].append(b)
            self.e[b].append(a)
        
        self.dfs(1)
        print(self.res)

        for e in reversed(edges):
            if e in self.res:
                return e

    def dfs(self, t, pre=None):
        if self.v[t]:
            if pre != None:
                self.res.append([min(t, pre), max(t, pre)])
            return t

        self.v[t] = True
        for nxt in self.e[t]:
            if nxt != pre:
                res = self.dfs(nxt, t)
                if res != None:
                    if res == t:
                        return -1
                    if res > 0 and pre != None :                            
                        self.res.append([min(t, pre), max(t, pre)])
                    return res
        return None
