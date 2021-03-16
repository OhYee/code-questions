func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
    for i:=0; i<n; i++ {
        matrix[i] = make([]int, n)
    }
  
    x := 0
    y := 0
    
    // 0   x 下
    // 1   x 上
    // 2   y 右
    // 3   y 左
    status := 2

    maxn := n
    maxm := n
    minn := 0
    minm := 0

    num := 1
    for num <= n * n {
        matrix[x][y] = num
        num++

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
    return matrix
}
