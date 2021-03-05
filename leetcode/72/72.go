func min3(a, b, c int) int {
    return min(a, min(b, c))
}

func min(a, b int) int {
    if a < b {
        return a
    }   
    return b 
}

func minDistance(word1 string, word2 string) int {
    n := len(word1)
    m := len(word2)

    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
        dp[i][0] = i
    }
    for j := 0; j <= m; j++ {
        dp[0][j] = j
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1] // 相当于直接添加新的字符，不需要任何变动
            } else {
                dp[i][j] = min3(
                    dp[i-1][j-1],   // 添加新的字符，并且随便变化一个
                    dp[i-1][j],     // 字符串 1 加一个字符
                    dp[i][j-1],     // 字符串 2 加一个字符
                ) + 1
            }
        }
    }

    // for i := 0; i <= n; i++ {
    //     fmt.Println(dp[i])
    // }
    // fmt.Println()

    return dp[n][m]
}
