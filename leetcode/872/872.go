/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    l1 := getLeaves(root1)
    l2 := getLeaves(root2)

    // fmt.Println(l1, l2)
    if len(l1) != len(l2) {
        return false
    }
    for i := 0; i < len(l1); i++ {
        if l1[i] != l2[i] {
            return false
        }
    }
    return true
}

func getLeaves(root *TreeNode) []int {
    leaves := make([]int, 0)
    stack := make([]*TreeNode, 0)

    stack = append(stack, root)
    for len(stack) != 0 {
        cur := stack[len(stack) - 1]
        stack = stack[0:len(stack) - 1]
        // fmt.Println(cur)
        if cur == nil {
            continue
        }
        if cur.Left == nil && cur.Right == nil {
            leaves = append(leaves, cur.Val)
        }
        if cur.Left != nil {
            stack = append(stack, cur.Left)
        }
        if cur.Right != nil {
            stack = append(stack, cur.Right)
        }
    }
    return leaves
}
