func removeDuplicates(S string) string {
    n := len(S)
    remove := make([]bool, n)

    for {
        flag := false
        for i:=0; i<n; i++ {
            if remove[i] {
                continue
            }
            for j:=i+1; j<n; j++ {
                if !remove[j] {
                    if S[i] == S[j] {
                        remove[i] = true
                        remove[j] = true
                        flag = true
                    }
                    break
                }
            }
            if flag {
                break
            }
        }
        if !flag {
            break
        }
    }

    sb := strings.Builder{}    
    for i:=0; i<n; i++ {
        if !remove[i] {
            sb.WriteByte(S[i])
        }
    }
    return sb.String()
}
