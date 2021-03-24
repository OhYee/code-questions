type Pair struct {
    Min int
    Max int
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func find132pattern(nums []int) bool {
    n := len(nums)
    
    stack := make([]int, n)
    top := 0

    leftMin := make([]int, n)
    leftMin[0] = math.MaxInt32
    for i := 1; i < n; i++ {
        leftMin[i] = min(leftMin[i-1], nums[i-1])
    }

    for i := n-1; i >= 0; i-- {
        num := nums[i]

        temp := math.MaxInt32
        for top != 0 && stack[top-1] < num {
            top--
            temp = stack[top]
        } 
        if leftMin[i] < temp && leftMin[i] < num && num > temp {
            return true
        }

        stack[top] = num
        top++
    }

    return false
}
