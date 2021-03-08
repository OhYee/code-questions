func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func stoneGame(piles []int) bool {
    return true
    n := len(piles)
    dp := make([][]int, n)
    for i:=0; i<n; i++ {
        dp[i] = make([]int, n)
    }

    for i:=0; i<n; i++ {
        for j:=i+1; j<n; j++ {
            fix := -1
            if (j - i) % 2 == 1 {
                fix = 1
            } 
            dp[i][j] = max(
                dp[i+1][j] + piles[i] * fix,
                dp[i][j-1] + piles[i] * fix,
            )
        }
    }

    return dp[0][n-1] > 0
}

