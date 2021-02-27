func longestSubstring(s string, k int) int {
    if len(s) == 0 {
        return 0
    }

    // 统计各个字符的数目
    count := make([]int, 26)
    for _, c := range s {
        count[c - 'a']++
    }

    pre := 0
    res := -1
    for i, c := range s {
        if (count[c - 'a'] < k) {
            // 使用数目不足的数进行分割，求各个区间的结果取最大值
            t := longestSubstring(s[pre:i], k)
            if t > res {
                res = t
            }
            pre = i + 1
        } 
    }
    if res == -1 {
        // 所有数都合法
        res = len(s)
    } else {
        // 处理最后一个区间
        t := longestSubstring(s[pre:], k)
        if t > res {
            res = t
        }
    }
    // fmt.Println(s,res)
    return res
}
