/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    newRoot := &TreeNode {
        Val: 0,
        Left: nil,
        Right: nil,
    }

    cur := newRoot
    var dfs func(t *TreeNode)
    dfs = func(t *TreeNode) {
        if t == nil {
            return
        }

        dfs(t.Left)

        cur.Right = &TreeNode {
            Val: t.Val,
            Left: nil,
            Right: nil,
        }
        cur = cur.Right
        
        dfs(t.Right)
    }

    dfs(root)
    return newRoot.Right
}
