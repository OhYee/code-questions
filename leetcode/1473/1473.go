func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

type Status [3]int
var IF = map[bool]int{true: 1, false: 0}

func minCost(houses []int, cost [][]int, m, n, target int) int {
    const INF = math.MaxInt32
    mem := make(map[Status]int)


    var dfs func (cur, target, color int) int
    dfs = func(cur, target, color int) (res int) {
        res = INF

        if r, e := mem[Status{cur, target, color}]; e {
            return r
        }
        defer func () {
            mem[Status{cur, target, color}] = res
        } ()

        if target == -1 || cur + target > m {
            return INF
        }
        if cur == m {
            return 0
        }

        if houses[cur] != 0 {
            res = min(
                res,
                dfs(
                    cur + 1, 
                    target - IF[houses[cur] != color], 
                    houses[cur],
                ),
            )
        } else {
            for i := 1; i <= n; i++ {
                res = min(
                    res,
                    dfs(
                        cur + 1, 
                        target - IF[i != color], 
                        i,
                    ) + cost[cur][i-1],
                )
            }
        }
        return
    }
    
    ans := dfs(0, target, -1)
    if ans == INF {
        ans = -1
    }

    return ans
}
