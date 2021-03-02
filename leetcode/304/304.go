type NumMatrix struct {
    sum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    n := len(matrix)
    if n == 0 {
        return NumMatrix {}
    }
    m := len(matrix[0])
    sum := make([][]int, n)

    sum[0] = make([]int, m)
    sum[0][0] = matrix[0][0]
    for j := 1; j < m; j++ {
        sum[0][j] = sum[0][j-1] + matrix[0][j]
    }
    for i := 1; i < n; i++ {
        sum[i] = make([]int, m)
        sum[i][0] = sum[i-1][0] + matrix[i][0]
        for j := 1; j < m; j++ {
            sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i][j]
        }
    }
    return NumMatrix {
        sum: sum,
    }
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    var part0, part1, part2, part3 int
    part0 = this.sum[row2][col2]
    if row1 != 0 {
        part1 = this.sum[row1-1][col2]
    }
    if col1 != 0 {
        part2 = this.sum[row2][col1-1]
    }
    if row1 != 0 && col1 != 0 {
        part3 = this.sum[row1-1][col1-1]
    }
    return  part0 - part1 - part2 + part3
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
