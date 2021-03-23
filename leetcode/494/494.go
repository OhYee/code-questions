func findTargetSumWays(nums []int, S int) int {
    n := len(nums)

    sm := 0
    for i:=0; i<n; i++ {
        sm += nums[i]
    }
    if S > sm {
        return 0
    }
    sm2 := 2*sm
    sm4 := 2*sm2

    dp := make([][]int, n)
    for i:=0; i<n; i++ {
        dp[i] = make([]int, sm4+1)
    }

    dp[0][sm2 + nums[0]] += 1
    dp[0][sm2 - nums[0]] += 1

    for i:=1; i<n; i++ {
        for j:=0; j<=sm4; j++ {
            if j-nums[i] >= 0 && j-nums[i] <= sm4 {
                dp[i][j] += dp[i-1][j-nums[i]]
            }
            if j+nums[i] >= 0 && j+nums[i] <= sm4 {
                dp[i][j] += dp[i-1][j+nums[i]]
            }
        }
    }

    // for i:=0; i<n; i++ {
    //     fmt.Println(dp[i])
    // }
    
    return dp[n-1][S+sm2]
}
