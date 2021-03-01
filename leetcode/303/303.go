type NumArray struct {
    sum []int
}


func Constructor(nums []int) NumArray {
    sum := make([]int, len(nums))
    for i, num := range nums {
        if i == 0 {
            sum[i] = num
        } else {
            sum[i] = sum[i-1] + num
        }
    }
    return NumArray {
        sum: sum,
    }
}


func (this *NumArray) SumRange(i int, j int) int {
    if i == 0 {
        return this.sum[j]
    }
    return this.sum[j] - this.sum[i - 1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
