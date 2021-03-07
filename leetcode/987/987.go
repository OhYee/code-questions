/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
    "sort"
)

type Obj struct {
    Val int
    Row int
    Col int
}

type Objs []Obj

func (objs Objs) Len() int {
    return len(objs)
}

func (objs Objs) Swap(i, j int) {
    objs[i], objs[j] = objs[j], objs[i]
}

func (objs Objs) Less(i, j int) bool {
    if objs[i].Col == objs[j].Col {
        if objs[i].Row == objs[j].Row {
            return objs[i].Val < objs[j].Val
        }
        return objs[i].Row < objs[j].Row
    }
    return objs[i].Col < objs[j].Col
}


func verticalTraversal(root *TreeNode) [][]int {

    values := make([]Obj, 0)
  
    var dfs func (root *TreeNode, row, col int)
    dfs = func (root *TreeNode, row, col int) {
        if root != nil {
            values = append(values, Obj{
                Val: root.Val,
                Row: row,
                Col: col,
            })
            dfs(root.Left, row+1, col-1)
            dfs(root.Right, row+1, col+1)
        }
    }
    dfs(root, 0, 0)

    sort.Sort(Objs(values))

    res := make([][]int, 0)
    pre := 9999999
    pos := -1
    for _, elem := range values {
        if elem.Col != pre {
            res = append(res, []int{})
            pos++
        }
        res[pos] = append(res[pos], elem.Val)
        pre = elem.Col
    }

    return res
}
