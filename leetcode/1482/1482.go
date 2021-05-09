func minDays(bloomDay []int, m int, k int) int {
    n := len(bloomDay)

    max := 0
    for i := 0; i < n; i++ {
        if bloomDay[i] > max {
            max = bloomDay[i]
        }
    }

    check := func(x int) (res bool) {
        // defer func () {
        //     fmt.Println("check", x, res)
        // }()

        cnt := 0
        tmp := 0
        for i := 0; i < n; i++ {
            if bloomDay[i] <= x {
                tmp++
                if tmp == k {
                    tmp = 0
                    cnt++
                    if cnt >= m {
                        return true
                    }
                }
            } else {
                tmp = 0
            }
        }
        return false
    }

    res := sort.Search(max + 1, check)
    if res > max {
        res = -1
    }

    // fmt.Println()
    return res
}
