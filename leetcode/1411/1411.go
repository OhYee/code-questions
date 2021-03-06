const m = 1e9 + 7

func numOfWays(n int) int {
    a, b := do(n)
    return (a+b)%m
}

func do(n int) (r2 int, r3 int) {
	if n == 0 {
		r2 = 0
		r3 = 0
		return
	}
	if n == 1 {
		r2 = 6
		r3 = 6
		return
	}

	n2, n3 := do(n - 1)
	r2 = ((n2*3)%m + (n3*2)%m)%m
	r3 = ((n2*2)%m + (n3*2)%m)%m
	return
}
