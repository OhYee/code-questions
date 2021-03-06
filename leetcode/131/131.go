func partition(s string) [][]string {
    n := len(s)
    dp := make([][]bool, n)
    for i:=0; i<n; i++ {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }
    for i:=n-1; i>=0; i-- {
        for j:=i+1; j<n; j++ {
            if s[i] == s[j] && (i+1>=j-1 || dp[i+1][j-1]) {
                dp[i][j] = true
            }
        }
    }

    var dfs func(a, l int) [][]string
    dfs = func (a, l int)[][]string {
        parts := make([][]string, 0, (1+l)*l/2)
        if dp[a][a+l-1] {
            parts = append(parts, []string{s[a:a+l]})
        }
        
        // 开始遍历分割
        for i:=1; i<l; i++ {
            // 分割成 a~a+i-1 和 a+i~a+l
            if dp[a][a+i-1] {
                // 左部分是回文串，合法
                leftPart := s[a:a+i]
                rightParts := dfs(a+i, l-i)
                for _, rightPart := range rightParts {
                    parts = append(parts, append([]string{leftPart}, rightPart...))
                }
            }
        }
        return parts
    }

    return dfs(0, n)
}


