package main

import (
	"sort"
)

func largestPerimeter(A []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for i := 2; i < len(A); i++ {
		if A[i-2] < A[i-1]+A[i] {
			return A[i-2] + A[i-1] + A[i]
		}
	}
	return 0
}
