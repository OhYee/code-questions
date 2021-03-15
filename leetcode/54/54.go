func spiralOrder(matrix [][]int) []int {
    x := 0
    y := 0

    n := len(matrix)
    m := len(matrix[0])
    
    // 0   x 下
    // 1   x 上
    // 2   y 右
    // 3   y 左
    status := 2

    maxn := n
    maxm := m
    minn := 0
    minm := 0

    res := make([]int, n*m)
    pos := 0
    for pos < m * n {
        res[pos] = matrix[x][y]
        pos++
        switch status {
            case 0:
                x++
            case 1:
                x--
            case 2:
                y++
            case 3:
                y--
        }
        if x >= maxn {
            x--
            y--
            status = 3
            maxm--
        }
        if y >= maxm {
            y--
            x++
            status = 0
            minn++
        }
        if x < minn {
            x++
            y++
            status = 2
            minm++
        }
        if y < minm {
            y++
            x--
            status = 1
            maxn--
        }
    }
    return res
}
