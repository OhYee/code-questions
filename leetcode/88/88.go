func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i:=m+n-1; i>=n; i-- {
        nums1[i] = nums1[i-n]
    }
    fmt.Println(nums1)
    l := n
    r := 0
    pos := 0
    for pos < n + m {
        if r >= n || (l < n + m && nums1[l] < nums2[r]) {
            nums1[pos] = nums1[l]
            l++
            pos++
        } else {
            nums1[pos] = nums2[r]
            r++
            pos++
        }
    }
}
