package main

func minPatches(nums []int, n int) int {
    var m, un, res, i, l = uint(1), uint(n), 0, 0, len(nums) 
    for m <= un {
        if i < l && uint(nums[i]) <= m {
            m += uint(nums[i])
            i++
        } else {
            m += m
            res++
        }
    }
    return res
}
