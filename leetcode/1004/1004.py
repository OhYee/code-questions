class Solution:
    def longestOnes(self, A: List[int], K: int) -> int:
        n = len(A)

        # 统计从 [0, i] 中 0 的个数，num0[-1] 为最后一个，始终为 0
        num0 = [0 for i in range(n+1)]
        for i in range(0, n):
            num0[i] = num0[i-1] + (1 if A[i] == 0 else 0)

        # print(num0)
        res = 0
        i = 0
        while i + res < n:
            # 已知 [i, i+res) 区间可行
            # print(i, i+res,res)
            # 尝试 [i, i+res+1) 区间
            if num0[i+res] - num0[i-1] <= K:
                # 新区间合法
                res += 1
            else:
                # 新区间不合法
                i += 1
        return res
