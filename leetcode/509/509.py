class Solution:
    def fib(self, n: int) -> int:
        if n == 0:
            return 0
        if n == 1:
            return 1
        ppre = 0
        pre = 1
        res = 0
        for i in range(2, n+1):
            res = ppre + pre
            ppre = pre
            pre = res
        return res
