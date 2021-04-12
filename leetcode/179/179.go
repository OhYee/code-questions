
func largestNumber(nums []int) string {
    sort.Slice(nums, func (i, j int) bool {
        a, b := nums[i], nums[j]
        aa, _ := strconv.ParseInt(fmt.Sprintf("%d%d", a, b), 10, 64)
        bb, _ := strconv.ParseInt(fmt.Sprintf("%d%d", b, a), 10, 64)
        return aa > bb
    })

    builder := strings.Builder{}
    for _, num := range nums {
        builder.WriteString(fmt.Sprintf("%d", num))
    }
    res := strings.TrimLeft(builder.String(), "0")
    if len(res) == 0 {
        res = "0"
    }
    return res
}
