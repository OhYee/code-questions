class Solution:
    def maxSatisfied(self, customers: List[int], grumpy: List[int], X: int) -> int:
        n = len(customers)

        # 预计算原本的满意人数
        res = 0
        for i in range(n):
            res += customers[i] * (grumpy[i] ^ 1)

        # 滑动窗口计算满意人数
        # 第一个区间预计算
        mx = 0
        temp = 0
        for i in range(X):
            temp += customers[i] * grumpy[i]
        # print(0, X, temp)
        mx = max(mx, temp)
        for i in range(1, n-X+1):
            # [i, i + X) 范围不生气
            temp = (
                temp - 
                customers[i-1] * grumpy[i-1] + 
                customers[i+X-1] * grumpy[i+X-1]
            )
            # print(i, i+X, temp)
            mx = max(mx, temp)
        return res + mx
