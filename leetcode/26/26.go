func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }

    i := 0
    pre := nums[0] - 1
    pos := 0
    for i < n {
        if pre != nums[i] {
            pos++
        }
        nums[pos-1] = nums[i]
        pre = nums[i]
        i++
    }
    return pos
}
