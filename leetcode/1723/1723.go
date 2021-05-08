func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func minimumTimeRequired(jobs []int, k int) int {
    n := len(jobs)

    res := math.MaxInt32
    vis := make(map[string]struct{})

    var dfs func (status []int, cur int, maxCost int)
    dfs = func (status []int, cur int, maxCost int) {
        if res < maxCost {
            return
        }
        if cur >= n {
            res = maxCost
            return
        }

        temp := make([]int, k)
        copy(temp, status)
        sort.Ints(temp)
        hash := fmt.Sprintf("%+v", temp)

        if _, e := vis[hash]; e {
            return
        }
        vis[hash] = struct{}{}

        for i := 0; i < k; i++ {
            status[i] += jobs[cur]
            dfs(status, cur + 1, max(maxCost, status[i]))
            status[i] -= jobs[cur]
        }
    }

    status := make([]int, k)
    dfs(status, 0, 0)

    return res
}
