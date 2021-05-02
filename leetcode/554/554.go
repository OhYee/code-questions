func leastBricks(wall [][]int) int {
    m := make(map[int]int)
    for _, line := range wall {
        t := 0
        for _, w := range line {
            t += w
            m[t]++
        }
    }
    first := 0
    second := 0
    for _, v := range m {
        if v > first {
            second = first
            first = v         
        } else if v > second {
            second = v
        }
    }
    // fmt.Println(first, second, m)
    return len(wall) - second
}
