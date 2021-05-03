func reverse(x int) int {
    neg := false
    if x < 0 {
        neg = true
        x = -x
    } 

    res := 0
    for x > 0 {
        res = res * 10 + (x % 10)
        x /= 10
    }
    if neg {
        res *= -1
    }
    if res > math.MaxInt32 || res < math.MinInt32 {
        res = 0
    }

    return res
}
