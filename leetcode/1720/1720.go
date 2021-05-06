func decode(encoded []int, first int) []int {
    n := len(encoded)
    arr := make([]int, n + 1)
    arr[0] = first
    for i := 1; i <= n; i++ {
        arr[i] = arr[i-1] ^ encoded[i-1]
    }
    return arr
}
