func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func deleteAndEarn(nums []int) int {
    cnt := make(map[int]int)
    for _, num := range nums {
        cnt[num]++
    }

    n := len(cnt)
    keys := make([]int, n)
    pos := 0
    for k := range cnt {
        keys[pos] = k
        pos++
    }
    sort.Ints(keys)

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, 2)
    }

    dp[0][0] = 0
    dp[0][1] = cnt[keys[0]] * keys[0]
    for i := 1; i < n; i++ {
        cur := keys[i]
        if keys[i] == keys[i-1] + 1 {
            // 相邻
            dp[i][0] = max(dp[i-1][0], dp[i-1][1])
            dp[i][1] = dp[i-1][0] + cur * cnt[cur]
        } else {
            // 不相邻
            dp[i][0] = max(dp[i-1][0], dp[i-1][1])
            dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + cur * cnt[cur]
        }
        
    }
    fmt.Println(dp)
    return max(dp[n-1][0], dp[n-1][1])
}
