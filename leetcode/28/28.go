func strStr(haystack string, needle string) int {
    l1 := len(haystack)
    l2 := len(needle)
    if l2 == 0 {
        return 0
    }
    if l1 == 0 {
        return -1
    }

    next := make([]int, l2)
    next[0] = -1
    for k := 1; k < l2; k++ {
        t := next[k - 1]
        for t != -1 && needle[k-1] != needle[t]  {
            t = next[t]
        }
        if t != -1 && needle[k-1] == needle[t] {
            next[k] = t + 1
        } else {
            next[k] = 0
        }
    }

    // fmt.Println(next)

    i := 0
    j := 0
    for i != l1 {
        // fmt.Printf("judge %d %c %d %c\n", i, haystack[i], j, needle[j]);
        for (j != -1 && haystack[i] != needle[j]) {
            j = next[j]
        }
        i++
        j++
        if j == l2 {
            // fmt.Printf("match at %d\n", i - l2);
            return i - l2
            // i--
            // j = next[j - 1]
        }
    }
    return -1
}
