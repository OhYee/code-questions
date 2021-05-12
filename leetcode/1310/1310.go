func xorQueries(arr []int, queries [][]int) []int {
    n := len(arr)

    xor := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        xor[i] = xor[i - 1] ^ arr[i - 1]
    }
    // fmt.Println(xor)

    m := len(queries)
    res := make([]int, m)
    for i, q := range queries {
        l, r := q[0], q[1]
        res[i] = xor[r + 1] ^ xor[l]
    }

    return res
}
