class Solution:
    def minKBitFlips(self, A: List[int], K: int) -> int:
        n = len(A)
        diff = [False for i in range(n+1)]

        # True  与前一个不一样
        # False 与前一个相同
        for i in range(n):
            if i == 0:
                diff[i] = (0 ^ A[i]) != 0
            else:
                diff[i] = (A[i] ^ A[i-1]) != 0

        res = 0
        ok = True
        pre = False
        for i in range(n):
            if i + K - 1 < n:
                if (pre ^ diff[i]) == False:
                    # 当前为 0
                    # 修改后续区间 K
                    res += 1
                    diff[i] = not diff[i]
                    diff[i+K] = not diff[i+K]
            else:
                if (pre ^ diff[i]) == False:
                    ok = False
            pre ^= diff[i]
        # print(A, diff, res, ok)
        return res if ok else -1
