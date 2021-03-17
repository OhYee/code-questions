
func check(s string) bool {
    n := len(s)
    for i:=0; i<n/2; i++ {
        if s[i] != s[n-i-1] {
            return false
        }
    }
    return true
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minCut(s string) int {
    n := len(s)

    // 回文串 dp
    hw := make([][]bool, n)
    for i:=0; i<n; i++ {
        hw[i] = make([]bool, n)
        hw[i][i] = true
    }
    for i:=n-1; i>=0; i-- {
        for j:=i+1; j<n; j++ {
            hw[i][j] = (s[i] == s[j]) && (hw[i+1][j-1] || i+1 > j-1)
        }
    }
    
    // 最少分割 dp
    dp := make([]int, n)
    for i:=1; i<n; i++ {
        dp[i] = dp[i-1] + 1
        // 找到以 i 结尾的最长回文串
        for j:=0; j<i; j++ {
            if hw[j][i] {
                if j==0 {
                    dp[i] = 0
                } else {
                    dp[i] = min(dp[i], dp[j-1] + 1)
                }
                // break
            }
        }
    }

    // for i:=0; i<n; i++ {
    //     fmt.Println(hw[i])
    // }
    // fmt.Println(dp)

    return dp[n-1]
}
