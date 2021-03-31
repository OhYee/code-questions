var res [][]int

func dfs(nums, now []int, cur int) {
    tmp := make([]int, len(now))
    copy(tmp, now)
    res = append(res, tmp)

    pre := -1000
    for i := cur; i < len(nums); i++ {
        if nums[i] == pre {
            continue
        }
        pre = nums[i]

        now = append(now, nums[i])
        dfs(nums, now, i + 1)
        now = now[:len(now)-1]
    }
}

func subsetsWithDup(nums []int) [][]int {
    if len(nums) == 0{
        return [][]int{{}}
    }
    sort.Ints(nums)
    res = make([][]int, 0)
    now := make([]int, 0, len(nums))

    dfs(nums, now, 0)

    return res
}
