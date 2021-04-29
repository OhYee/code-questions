
fun canCross(stones []int) bool {
    n := len(stones)

    dp := make([]bool, n)
    vis := make(map[[2]int]struct{})

    var jump func (pos int, from int)
    jump = func (pos int, from int) {
        if pos <= from {
            return
        }
        if _, e := vis[[2]int{pos, from}]; e {
            return
        }

        idx := sort.SearchInts(stones, pos)
        // fmt.Println(pos, from, stones)

        if idx < n && stones[idx] == pos {
            dp[idx] = true
            vis[[2]int{pos, from}] = struct{}{}

            d := pos - from
            jump(pos + d - 1, pos)
            jump(pos + d, pos)
            jump(pos + d + 1, pos)
        }
    }

    if stones[1] - stones[0] == 1 {
        jump(stones[1], stones[0])
        return dp[n-1]
    } 
    return false
}
