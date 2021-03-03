// import (
//     "math/bits"
// )

func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 0; i <= num; i++ {
        // res[i] = bits.OnesCount(uint(i))    
        if i & 1 == 0 {
            // 偶数
            res[i] = res[i/2]
        } else {
            // 奇数
            res[i] = res[i-1] + 1
        }
    }
    return res
}
