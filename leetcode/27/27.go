func removeElement(nums []int, val int) int {
    n := len(nums)
    pos := 0
    for i := 0; i < n; i++ {
        if (nums[i] != val) {
            nums[pos] = nums[i]
            pos++
        }
    }
    return pos
}
