class Solution:
    def maxProfit(self, k: int, prices: List[int]) -> int:
        n = len(prices)
        if n == 0:
            return 0

        
        # dp[i][j][0|1]
        # 第 i 天使用 j 次机会时（当前持有时）能达到的最大利润
        dp = [
            [
                [
                    -99999999 
                    for ii in range(2)
                ] for j in range(k + 1)
            ] for i in range(2)
        ]
        dp[0][0][0] = 0
        dp[0][0][1] = -prices[0]

        for i in range(1, n):
            for j in range(k + 1):
                t = i & 1
                p = t ^ 1
                dp[t][j][0] = max(dp[p][j][0], dp[p][j-1][1] + prices[i] if j>0 else 0) 
                dp[t][j][1] = max(dp[p][j][0] - prices[i], dp[p][j][1])
        res = 0
        t = (n-1) & 1
        for i in range(k + 1):
            res = max(res, dp[t][i][0], dp[t][i][1]) 
        return res
