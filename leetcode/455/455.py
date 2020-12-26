class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        n = len(g)
        g.sort()
        s.sort()
        res = 0
        for c in s:
            if res < n and g[res] <= c:
                res += 1
            if res >= n:
                break
        return res            
