func searchMatrix(matrix [][]int, target int) bool {
    n := len(matrix)
    m := len(matrix[0])

    idx := sort.Search(n * m, func (i int) bool {
        return matrix[i / m][i % m] >= target
    })
    return idx < n * m && matrix[idx / m][idx % m] == target
}
