func findKthPositive(arr []int, k int) int {
    n := len(arr)

    count := 0
    num := 1 // arr[0] + 1
    for count < k {
        if num <= arr[n-1] {
            idx := sort.SearchInts(arr, num)
            if idx < n && arr[idx] == num {
                num++
                continue
            }
            count++
            num++
        } else {
            num += k - count
            count = k
            break
        }
    }

	return num - 1
}
