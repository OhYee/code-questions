func numRabbits(answers []int) int {
    count := make(map[int]int)
    for _, ans := range answers {
        count[ans]++
    }

    // fmt.Println(count)

    res := 0
    for k, v := range count {
        k++
        if v > k {
            groups := v / k
            if v % k != 0 {
                groups++
            }
            res += k * groups  
        } else {
            // 计数 小于等于
            res += k
        }
        // fmt.Println(k,v,res)
    }
    return res
}
