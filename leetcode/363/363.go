func maxSumSubmatrix(matrix [][]int, k int) int {
    n := len(matrix)
    m := len(matrix[0])

    sum := make([][]int, n)
    for i := 0; i < n; i++ {
        sum[i] = make([]int, m)
        for j := 0; j < m; j++ {
            sum[i][j] = matrix[i][j]
            if i > 0 {
                sum[i][j] += sum[i-1][j]
            }
            if j > 0 {
                sum[i][j] += sum[i][j-1]
            }
            if i > 0 && j > 0 {
                sum[i][j] -= sum[i-1][j-1]
            }
        }
    }

    res := -math.MaxInt32
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            for p:=i; p<n; p++ {
                for q:=j; q<m; q++ {
                    s := sum[p][q]
                    if i > 0 {
                        s -= sum[i-1][q]
                    }
                    if j > 0 {
                        s -= sum[p][j-1]
                    }
                    if i > 0 && j > 0 {
                        s += sum[i-1][j-1]
                    }

                    if s <= k && s > res {
                        res = s
                    }
                }
            }
        }
    }
    return res
}
