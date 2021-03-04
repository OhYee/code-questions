class Solution:
    def maxEnvelopes(self, envelopes: List[List[int]]) -> int:      
        n = len(envelopes)
        envelopes.sort(key=lambda x: (x[0], -x[1]))

        stack = [0] * n
        pos = 0
        for i in range(n):
            t = envelopes[i][1]
            if pos == 0 or stack[pos-1] < t:
                stack[pos] = t
                pos += 1
            else:
                p = bisect.bisect(stack[:pos], t)
                if p == 0 or stack[p-1] != t:
                    stack[p] = t
        return pos
