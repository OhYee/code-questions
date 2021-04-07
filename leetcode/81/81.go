func BinarySearch(nums []int, l, r, target int) int {
    if l == r && nums[l] == target {
        return l
    } else if l > r || ( l == r && nums[l] != target)  {
        return -1
    }

    mid := (l + r) / 2
    // 判断区间类型
    if nums[l] >= nums[r] {
        // 旋转区间
        if nums[mid] == target {
            return mid
        }
        // 两侧区间都需要搜索
        if pos := BinarySearch(nums, l, mid-1, target); pos != -1 {
            return pos
        }
        return BinarySearch(nums, mid+1, r, target)
    } else {
        // 有序区间
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            // 左侧
            return BinarySearch(nums, l, mid-1, target)
        } else {
            return BinarySearch(nums, mid+1, r, target)
        }
    }
}

func search(nums []int, target int) bool {
    return BinarySearch(nums, 0, len(nums)-1, target) != -1
}
