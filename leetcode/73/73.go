func setZeroes(matrix [][]int)  {
    n := len(matrix)
    m := len(matrix[0])

    r0 := false
    for i := 0; i < n; i++ {
        if matrix[i][0] == 0 {
            r0 = true
        }
    }
    c0 := false
    for j := 0; j < m; j++ {
        if matrix[0][j] == 0 {
            c0 = true
        }
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][j] == 0 {
                // 当前行和列需要置 0 
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    for i := 1; i < n; i++ {
        if matrix[i][0] == 0 {
            for j := 1; j < m; j++ {
                matrix[i][j] = 0
            }
        }
    }

    for j := 1; j < m; j++ {
        if matrix[0][j] == 0 {
            for i := 1; i < n; i++ {
                matrix[i][j] = 0
            }
        }
    }

    if r0 {
        for i := 0; i < n; i++ {
            matrix[i][0] = 0
        }
    }

    if c0 {
        for j := 0; j < m; j++ {
            matrix[0][j] = 0
        }
    }
    
}
