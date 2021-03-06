
func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n == 0{
        return 0
    }
    count := make(map[byte]int)
    i := 0
    l := 1
    count[s[0]] = 1
    morethan1 := 0
    for i+l < n {
        // fmt.Println(s[i:i+l], count, morethan1)
        count[s[i+l]]++
        c := count[s[i+l]]
        if c == 2 {
            // 超出限制
            morethan1++
        }

        if morethan1 > 0 {
            // 当前不合法
            count[s[i]]--
            if count[s[i]] == 1 {
                morethan1--
            }
            i++
        } else {
            // 合法
            l++
        }
    }
    return l
}
