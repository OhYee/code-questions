func max(a, b int) int {
    if a > b{
        return a
    }
    return b
}

func maxScoreSightseeingPair(values []int) int {
    n := len(values)
    /* 
        dp[i] = max{
            v[i] + v[i-1] - 1,
            v[i] + v[i-2] - 2,
            ……
            v[i] + v[0] - i,
        }
        dp[i-1] = max{
            v[i-1] + v[i-2] - 1,
            v[i-1] + v[i-3] - 2,
            ……
            v[i-1] + v[0] - (i-1),
        }
        dp[i] = max{
            v[i] + v[i-1] - 1,
            dp[i-1] - v[i-1] + v[i] - 1
        }
    */
    dp := make([]int, n)
    dp[1] = values[1] + values[0] - 1
    res := dp[1]
    for i:=2; i<n; i++ {
        dp[i] = max(
            values[i] + values[i-1] - 1,
            dp[i-1] - 1 - values[i-1] + values[i],
        )
        res = max(res, dp[i])
    }
    return res
}
