class Solution:
    def maxTurbulenceSize(self, arr: List[int]) -> int:
        n = len(arr)
        if n == 1:
            return 1
        def calc(a, b):
            if a == b:
                return 0
            elif a > b:
                return -1
            else:
                return 1
        res = 0
        count = 0
        pre = 0
        for i in range(1, n):
            t = calc(arr[i-1], arr[i])
            if t * pre == -1:
                count += 1
            else:
                res = max(count, res)
                count = 1 if t != 0 else 0
            pre = t
        res = max(count, res)
        return res + 1
