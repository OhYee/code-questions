package main

type pair struct{ f, t int }

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	nodeID := make(map[string]int)
	count := 0

	dis := make(map[pair]float64)
	for i, e := range equations {
		a, b, v := e[0], e[1], values[i]
		ia, ea := nodeID[a]
		if !ea {
			ia = count
			nodeID[a] = ia
			count++
		}
		ib, eb := nodeID[b]
		if !eb {
			ib = count
			nodeID[b] = ib
			count++
		}
		dis[pair{ia, ib}] = v
		dis[pair{ib, ia}] = 1 / v
	}

	for k := 0; k < count; k++ {
		for i := 0; i < count; i++ {
			for j := 0; j < count; j++ {
				if i == j {
					dis[pair{i, i}] = 1
				} else {
					_, e0 := dis[pair{i, j}]
					if !e0 {
						d1, e1 := dis[pair{i, k}]
						d2, e2 := dis[pair{k, j}]
						if e1 && e2 {
							dis[pair{i, j}] = d1 * d2
						}
					}
				}
			}
		}
	}

	res := make([]float64, len(queries))
	for i, query := range queries {
		var exist bool
		si, se := nodeID[query[0]]
		ei, ee := nodeID[query[1]]
		res[i], exist = dis[pair{si, ei}]
		if !exist || !se || !ee {
			res[i] = -1.0
		}
	}
	return res
}
