func isScramble(s1, s2 string) bool {
    n1 := len(s1)
    n2 := len(s2)
    if n1 != n2 {
        return false
    }

    dp := make([][][]int8, n1)
    for i := 0; i < n1; i++ {
        dp[i] = make([][]int8, n2)
        for j := 0; j < n2; j++ {
            dp[i][j] = make([]int8, n1 - i + 1)
            for k := 0; k < n1 - i + 1; k++ {
                dp[i][j][k] = -1
            }
        }
    }

    count := [26]int{}
    
    var check func(i1, i2, l int) int8
    check = func(i1, i2, l int) int8 {
        if dp[i1][i2][l] == -1 {
            // 字符相同
            x, y := s1[i1:i1+l], s2[i2:i2+l]
            if x == y {
                dp[i1][i2][l] = 1
            } 

            if dp[i1][i2][l] == -1 {
                // 字符频率不同
                for i := 0; i < 26; i++ {
                    count[i] = 0
                }
                for i, ch := range x {
                    count[ch-'a']++
                    count[y[i]-'a']--
                }
                for _, f := range count[:] {
                    if f != 0 {
                        dp[i1][i2][l] = 0
                    }
                }
            }

            if dp[i1][i2][l] == -1 {
                // 枚举分割位置
                for i := 1; i < l; i++ {
                    // 不交换的情况
                    if check(i1, i2, i) == 1 && check(i1+i, i2+i, l-i) == 1 {
                        dp[i1][i2][l] = 1
                    }
                    // 交换的情况
                    if check(i1, i2+l-i, i) == 1 && check(i1+i, i2, l-i) == 1 {
                        dp[i1][i2][l] = 1
                    }
                }
            }

            if dp[i1][i2][l] == -1 {
                // 检查失败,无法转移
                dp[i1][i2][l] = 0
            }
        }
        return dp[i1][i2][l]
    }
    return check(0, 0, n1) == 1
}
