
import (
	"container/list"
)

type Obj struct {
	Station int
	Dis     int
}

func getF(x int, f map[int]int) int {
	_, exist := f[x]
	if !exist {
		f[x] = x
	}
	if x == f[x] {
		return x
	}
	f[x] = getF(f[x], f)
	return f[x]
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	f := make(map[int]int)

	connect := func(x, y int) {
		xx := getF(x, f)
		yy := getF(y, f)
		f[xx] = yy
	}

	// 每个站点对应的线路列表
	m := make(map[int][]int)
	for idx, route := range routes {
		hasSource := false
		hasTarget := false
		pre := -1
		for _, s := range route {
			if s == source {
				hasSource = true
			}
			if s == target {
				hasTarget = true
			}

			if pre != -1 {
				connect(pre, s)
			}
			pre = s

			_, exist := m[s]
			if exist {
				m[s] = append(m[s], idx)
			} else {
				m[s] = []int{idx}
			}
		}
		if hasSource && hasTarget {
			return 1
		}
	}

	if getF(source, f) != getF(target, f) {
		return -1
	}

	vis := make(map[int]struct{})
	q := list.New()
	q.PushBack(Obj{
		Station: source,
		Dis:     0,
	})

	for k, v := range m {
		if len(v) <= 1 {
			// 站点无法到达其他线路，可以直接忽略
			vis[k] = struct{}{}
		}
	}

	for q.Len() != 0 {
		head := q.Front()
		q.Remove(head)

		obj := head.Value.(Obj)
		t := obj.Station
		d := obj.Dis

		// 当前点未到达过
		// 检查所有可以从这个点前往的点
		for _, rid := range m[t] {
			// 当前点所在的所有线路
			for _, s := range routes[rid] {
				// 线路所有站点
				if s == target {
					return d + 1
				}

				_, visited := vis[s]
				if !visited {
					q.PushBack(Obj{
						Station: s,
						Dis:     d + 1,
					})
					vis[s] = struct{}{}
				}
			}
		}
	}
	return -1
}

