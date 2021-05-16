func findMaximumXOR(nums []int) int {
    high := 30
    ans := 0

    for k := high; k >= 0; k-- {
    
        m := make(map[int]struct{})
        for _, num := range nums {
            m[num >> k] = struct{}{}
        }
    
        ans = (ans << 1) + 1


        all_not_in := true
        for num := range m {
            if _, e := m[num ^ ans]; e {
                all_not_in = false
                break
            }
        }

        if all_not_in {
            ans -= 1
        }

        // fmt.Println(m, ans)
    }

    return ans
}
