func judgeSquareSum(c int) bool {
    n := int(math.Sqrt(float64(c))) + 1
    for a := 0; a <= n; a++ {
        b := int(math.Sqrt(float64(c - a * a )))
        if a * a + b * b == c {
            return true
        }
    }
    return false
}
