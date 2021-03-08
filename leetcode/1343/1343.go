func numOfSubarrays(arr []int, k int, threshold int) int {
    n := len(arr)
    if n < k {
        return 0
    }

    target := k * threshold
    res := 0
    sum := 0
    i := 0
    for i=0; i<k; i++ {
        sum += arr[i]
    }
    i = 0
    // fmt.Println(arr[i:i+k], sum)
    if sum >= target {
        res++
    }
    for i + k < n {
        sum = sum - arr[i] + arr[i + k]
        i++
        // fmt.Println(arr[i:i+k], sum)
        if sum >= target {
            res++
        }
    }
    return res
}
