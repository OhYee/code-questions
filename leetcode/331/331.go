func New(cur int, s []string) (newCur int, res bool) {
    if cur >= len(s) {
        res = false
        return 
    }

    fmt.Printf("%s %d\n", s[cur], cur)
    newCur = cur + 1

    if s[cur] == "#" {
        res = true
        return 
    } else {
        if newCur, res = New(newCur, s); !res {
            return
        }
        if newCur, res = New(newCur, s); !res {
            return
        }
    }
    return
}

func isValidSerialization(preorder string) bool {
    lst := strings.Split(preorder, ",")
    cur, res := New(0, lst)    
    return res && cur == len(lst)
}
