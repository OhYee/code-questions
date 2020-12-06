class Solution:
    def concatenatedBinary(self, n: int) -> int:
        mod = 1e9+7
        res = 0
        for i in range(n+1):
            res = int((((res * (2**len(bin(i)[2:]))) % mod) + i) % mod)
        return res
