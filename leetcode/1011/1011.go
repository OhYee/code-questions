func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func check(weights []int, D int, pre int) bool {
    t := 1
    today := 0
    for _, w := range weights {
        if today + w <= pre {
            today += w
        } else {
            t += 1
            today = w
        }
    }
    return t <= D
}

func shipWithinDays(weights []int, D int) int {
    sum := 0
    mx := 0
    for _, w := range weights {
        sum += w
        mx = max(mx, w)
    }

    pre := max(sum / D, mx)
    return pre + sort.Search(sum-pre, func (i int) bool {
        return check(weights, D, pre + i)
    })
}
