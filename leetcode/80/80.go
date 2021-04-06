func removeDuplicates(nums []int) int {
    n := len(nums)
    pre :=  nums[0]
    pos := 1
    count := 1
    for i := 1; i < n; i++ {
        num := nums[i]
        if num != pre || (num == pre && count < 2) {
            nums[pos] = num
            pos++
        }
        if num == pre {
            count++
        } else {
            count = 1
        }
        pre = num
    }
    return pos
}
