/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    nums []int
    l int
    cur int
}


func dfs(this *TreeNode) []int {
    if this == nil {
        return []int{}
    }

    left := dfs(this.Left)
    right := dfs(this.Right)

    left = append(left, this.Val)
    left = append(left, right...)
    return left
}

func Constructor(root *TreeNode) BSTIterator {
    nums := dfs(root)
    return BSTIterator {
        nums: nums,
        l: len(nums),
        cur: 0,
    }
}

func (this *BSTIterator) Next() int {
    res := this.nums[this.cur]
    this.cur++
    return res
}


func (this *BSTIterator) HasNext() bool {
    return this.cur < this.l
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
