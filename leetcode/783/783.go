/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var pre = -math.MaxInt32
var res = math.MaxInt32

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func dfs(root *TreeNode) {
    if root == nil {
        return
    }
    dfs(root.Left)
    // fmt.Printf("(%d)%d ",pre, root.Val)
    res = min(res, root.Val-pre)
    pre = root.Val
    dfs(root.Right)
}

func minDiffInBST(root *TreeNode) int {
    pre = -math.MaxInt32
    res = math.MaxInt32
    dfs(root)
    // fmt.Println()
    return res
}
