import (
    "container/list"
)

func canReach(arr []int, start int) bool {
    if arr[start] == 0 {
        return true
    }

    n := len(arr)
    vis := make([]bool, n)

    q := list.New()
    q.PushBack(start)
    vis[start] = true
    for q.Len() != 0 {
        head := q.Front()
        q.Remove(head)
        t := head.Value.(int)

        // fmt.Println(t, arr[t], t+arr[t], t-arr[t])

        tt := t + arr[t]
        if tt >= 0 && tt < n && !vis[tt] {
            if arr[tt] == 0 {
                return true
            }
            q.PushBack(tt)
            vis[tt] = true
            
        } 

        tt = t - arr[t]
        if tt >= 0 && tt < n && !vis[tt] {
            if arr[tt] == 0 {
                return true
            }
            q.PushBack(tt)
            vis[tt] = true
        } 
    }
    return false
}
