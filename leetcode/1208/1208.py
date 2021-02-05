class Solution:
    def equalSubstring(self, s: str, t: str, maxCost: int) -> int:
        n = len(s)
        cost = [abs(ord(s[i]) - ord(t[i])) for i in range(n)]
        sums = [0 for i in range(n+1)]
        for i in range(1, n+1):
            sums[i] = sums[i-1] + cost[i-1]
        l = 1
        r = 0
        while r <= n:
            while r + 1 <= n and sums[r + 1] - sums[l - 1] <= maxCost:
                r += 1
            l += 1
            r += 1
        return r - l + 1

