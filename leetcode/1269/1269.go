const mod = 1e9+7

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func numWays(steps int, arrLen int) int {
    n := min(steps / 2 + 1, arrLen)
    
    cur := make([]int, n)
    cur[0] = 1
    for i := 0; i < steps; i++ {
        nxt := make([]int, n)
        for j := 0; j < n; j++ {
            if cur[j] > 0 {
                if j - 1 >= 0 {
                    nxt[j-1] = (nxt[j-1] + cur[j]) % mod
                }
                nxt[j] = (nxt[j] + cur[j]) % mod
                if j + 1 < n {
                    nxt[j+1] = (nxt[j+1] + cur[j]) % mod
                }
            }
        }
        cur = nxt
        // fmt.Println(cur)
    }
    return cur[0] % mod
}
