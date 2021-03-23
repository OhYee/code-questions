
export PATtedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    nums []int
    size int
    cur  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    nums := dfs(nestedList)
    return &NestedIterator {
        nums: nums,
        size: len(nums),
        cur:  0,
    }
}

func dfs(nestedList []*NestedInteger) []int {
    res := make([]int, 0, len(nestedList))
    for _, item := range nestedList {
        if item.IsInteger() {
            res = append(res, item.GetInteger())
        } else {
            res = append(res, dfs(item.GetList())...)
        }
    }
    return res
}

func (this *NestedIterator) Next() int {
    res := this.nums[this.cur]
    this.cur++
    return res
}

func (this *NestedIterator) HasNext() bool {
    return this.cur != this.size
}
