func BinaryFindMin(l, r int, nums []int) int {
    if (l == r) {
        return l
    } else if (l > r) {
        return -1
    }

    // 判断区间类型
    if (nums[l] < nums[r]) {
        // 递增类型
        return l
    } else {
        // 旋转类型
        mid := (l + r) / 2
        t := mid
        t2 := BinaryFindMin(l, mid-1, nums)
        if t2 != -1 && nums[t] > nums[t2] {
            t = t2
        }

        t2 = BinaryFindMin(mid+1, r, nums)
        if t2 != -1 && nums[t] > nums[t2] {
            t = t2
        }
        return t
    }
}

func findMin(nums []int) int {
    return nums[BinaryFindMin(0, len(nums)-1, nums)]
}
