func isMatch(s string, p string) bool {
    n := len(p)
    m := len(s)

    dp := make([][]bool, n+1)
    for i:=0; i<=n; i++{
        dp[i] = make([]bool, m+1)
    }
    dp[0][0] = true
    for i:=1; i<=n; i++ {
        if p[i-1] == '*'{
            dp[i][0] = dp[i-1][0]
        }
    }

    for i:=1; i<=n; i++ {
        for j:=1; j<=m; j++ {
            switch p[i-1] {
                case '?', s[j-1]:
                    dp[i][j] = dp[i-1][j-1]
                case '*':
                    dp[i][j] = dp[i-1][j-1] || dp[i-1][j] || dp[i][j-1]
                default:
                    dp[i][j] = false
            }
        }
    }

    // for i:=0; i<=n; i++ {
    //     fmt.Println(dp[i])
    // }
    // fmt.Println()

    return dp[n][m]
}
