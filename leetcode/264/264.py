type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func nthUglyNumber(n int) int {
    vis := make(map[int]struct{})
    h := &IntHeap{1}
    heap.Init(h)
    t := 0
    count := 0
    for {
        count++
        t = heap.Pop(h).(int)
        fmt.Println(t)
        if count == n {
            break
        }

        if _, e := vis[t*2]; !e {
            heap.Push(h, t*2)
            vis[t*2] = struct{}{}
        }
        if _, e := vis[t*3]; !e {
            heap.Push(h, t*3)
            vis[t*3] = struct{}{}
        }
        if _, e := vis[t*5]; !e {
            heap.Push(h, t*5)
            vis[t*5] = struct{}{}
        }
    } 
    return t
}
