func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    if n == 2 {
        return max(nums[0], nums[1])
    }
    if n == 3 {
        return max(nums[0], max(nums[1], nums[2]))
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, 2)
    }

    // 不偷第一家    
    dp[0][0] = 0
    dp[0][1] = 0
    for i := 1; i < n; i++ {
        // 不偷
        dp[i][0] = max(dp[i-1][0], dp[i-1][1])
        // 偷
        dp[i][1] = dp[i-1][0] + nums[i]
    }
    res := max(dp[n-1][0], dp[n-1][1])

    // 偷第一家
    dp[0][0] = 0
    dp[0][1] = nums[0]
    for i := 1; i < n; i++ {
        // 不偷
        dp[i][0] = max(dp[i-1][0], dp[i-1][1])
        // 偷
        dp[i][1] = dp[i-1][0] + nums[i]
    }
    res = max(res, dp[n-1][0])

    return res
}
