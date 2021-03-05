func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func trap(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }

    leftMax := make([]int, n)
    rightMax := make([]int, n)
    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i-1])
        rightMax[n-i-1] = max(rightMax[n-i], height[n-i])
    }

    res := 0
    for i := 1; i < n-1; i++ {
        h := height[i]
        lm := leftMax[i]
        rm := rightMax[i]
        hh := min(lm, rm)
        // fmt.Println(i, h, lm, rm, hh - h)
        if hh > h {
            res += hh - h
        }
    }
    return res
}
