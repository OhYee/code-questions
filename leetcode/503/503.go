import (
    "sort"
)

func nextGreaterElements(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return []int{}
    }

    // 先找到最大值的位置
    maxPos := 0
    for i:=1; i<n; i++ {
        if nums[maxPos] < nums[i] {
            maxPos = i
        }
    }

    // 栈需要保证单调递减
    // 新的数必定更接近于在判断的数，如果新的数更大，那么老数没有意义
    stack := make([]int, n)
    stackLen := 0

    stack[stackLen] = nums[maxPos]
    stackLen++

    res := make([]int, n)
    res[maxPos] = -1
    for i:=0; i<n; i++ {
        pos := (maxPos + n - 1 - i) % n 
        num := nums[pos]

        // 找到栈中第一个小于等于 num 的数（其左侧就是第一个的大于的数）
        lstPos := sort.Search(stackLen, func (p int) bool {
            return stack[p] <= num
        })

        // fmt.Println("find", num, "in", stack[:stackLen], lstPos)

        if lstPos == 0 {
            // 没有更大的数了
            res[pos] = -1
        } else {
            res[pos] = stack[lstPos-1]
        }
        // 栈内所有比当前数小的都可移除掉
        stack[lstPos] = num
        stackLen = lstPos + 1
    }
    return res
}
