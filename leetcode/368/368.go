func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func largestDivisibleSubset(nums []int) []int {
    n := len(nums)
    if n <= 1 {
        return nums
    }

    sort.Ints(nums)

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = []int{nums[i]}
    }

    idx := -1
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if (nums[i] % nums[j]) == 0 && len(dp[j]) + 1 > len(dp[i]) {
                copy(dp[i], dp[j])
                dp[i] = append(dp[i], nums[i])
            }
            if idx == -1 || len(dp[idx]) < len(dp[i]) {
                idx = i
            }
        }
    }

    fmt.Println(dp)
    return dp[idx]
}
