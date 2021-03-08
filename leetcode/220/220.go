type Obj struct {
    Val int
    Idx int
}

type Objs []Obj

func (o Objs) Less(i, j int) bool {
    if o[i].Idx == o[j].Idx {
        return o[i].Val < o[j].Val
    }
    return o[i].Idx < o[j].Idx
}

func (o Objs) Swap(i, j int) {
    o[i], o[j] = o[j], o[i]
}

func (o Objs) Len() int {
    return len(o)
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    var a int = 2147483647
    var b int = -1
    fmt.Println(a-b, abs(a-b))
    n := len(nums)
    objs := make([]Obj, n)
    for i:=0; i<n; i++ {
        objs[i] = Obj{
            Val: nums[i],
            Idx: i,
        }
    }

    sort.Sort(Objs(objs))

    for i:=0; i<n; i++ {
        for j:=i+1; j<n; j++ {
            if abs(objs[i].Idx - objs[j].Idx) > k {
                break
            }
            if abs(objs[i].Val - objs[j].Val) <= t {
                // fmt.Println(objs[i], objs[j])
                return true
            }
        }
    }
    return false
}
