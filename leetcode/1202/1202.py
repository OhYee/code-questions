class Solution:
    def smallestStringWithSwaps(self, s: str, pairs: List[List[int]]) -> str:
        n = len(s)
        f = [i for i in range(n)]
        def getF(x):
            f[x] = x if x == f[x] else getF(f[x])
            return f[x]

        # O(m)
        for pair in pairs:
            a, b = pair
            ra, rb = getF(a), getF(b)
            f[ra] = f[rb] = rb
        # print(f)
        
        # O(n)
        group = {}
        for i in range(n):
            root = getF(i)
            group[root] = group.get(root, "") + s[i]
        
        
        for k in group.keys():
            group[k] = sorted(group[k])

        # print(group)

        res = ""
        pos = {}
        for i in range(n):
            root = getF(i)
            ss = group[root]
            p = pos.get(root, 0)
            res += ss[p]
            pos[root] = p + 1

        return res
         
                
