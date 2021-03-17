func numDistinct(s string, t string) int {
    n := len(s)
    m := len(t)

    dp := make([][]int, (n+1)*(m+1))
    for i:=0; i<=n; i++ {
        dp[i] = make([]int, m+1)
        dp[i][0] = 1
    }

    for i:=1; i<=n; i++ {
        for j:=1; j<=m; j++ {
            if s[i-1] == t[j-1] {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }

    // for i:=0; i<=n; i++ {
    //     fmt.Println(dp[i])
    // }

    return dp[n][m]
}
