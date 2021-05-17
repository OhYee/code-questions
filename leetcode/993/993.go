/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    var px, py *TreeNode
    var dx, dy int

    var dfs func (*TreeNode, *TreeNode, int)
    dfs = func (parent *TreeNode, cur *TreeNode, depth int) {
        if cur == nil {
            return
        }
        if cur.Val == x {
            px = parent
            dx = depth
        }
        if cur.Val == y {
            py = parent
            dy = depth
        }
        dfs(cur, cur.Left, depth + 1)
        dfs(cur, cur.Right, depth + 1)
    }

    dfs(nil, root, 0)

    return px != py && dx == dy
}
