func isMatch(s string, p string) bool {
    n := len(p)
    m := len(s)
    dp := make([][]bool, n+1)
    for i:=0; i<=n; i++ {
        dp[i] = make([]bool, m+1)
        if i > 0 && p[i-1] == '*' {
            dp[i][0] = dp[i-2][0]
        }
    }
    dp[0][0] = true
    for i:=1; i<=n; i++ {
        if p[i-1] == '*' {
            dp[i][0] = dp[i-2][0]
        }
    }

    for i:=1; i<=n; i++ {
        for j:=1; j<=m; j++ {
            switch p[i-1] {
                case '.', s[j-1]:
                    // 字符相等
                    // 需要没有新字符前就相等
                    dp[i][j] = dp[i-1][j-1]
                case '*':
                    if p[i-2] == s[j-1] || p[i-2] == '.' {
                        // 新的字符通过 * 匹配成功
                        // 需要少匹配一次时相等 或 匹配 0 次时相等
                        dp[i][j] = dp[i][j-1] || dp[i-2][j]
                    } else {
                        // 匹配失败
                        // 匹配 0 次
                        dp[i][j] = dp[i-2][j]
                    }
                default:
                    dp[i][j] = false
            }
        }
    }

    return dp[n][m]

}

