func minimumDeviation(nums []int) int {
    root := NewTree()
    for _, num := range nums {
        if num & 1 == 1{
            root = root.Insert(num << 1)
        } else {
            root = root.Insert(num)
        }
    }

    max, _ := root.Max()
    min, _ := root.Min()
    res := max - min
    for {
        if max & 1 == 1 {
            break
        }
        root, _ = root.Remove(max)
        root = root.Insert(max >> 1)
        max, _ = root.Max()
        min, _ = root.Min()
        if max - min < res {
            res = max - min
        }
    }
    return res
}

// Tree 平衡二叉树结构
type Tree struct {
        number int
        count  int
        left   *Tree
        right  *Tree
        depth  int
}

// NewTree 生成一个平衡二叉树
func NewTree() *Tree {
        return nil
}

// NewTreeNode 新建一个二叉树节点
func NewTreeNode(number int) *Tree {
        return &Tree{
                number: number,
                count:  1,
                left:   nil,
                right:  nil,
                depth:  1,
        }
}

// Insert 向树内插入一个数
func (t *Tree) Insert(n int) *Tree {
        if t == nil {
                return NewTreeNode(n)
        }
        if n == t.number {
                // 就是当前值
                t.count++
                t.updateNode()
                return t
        } else if n > t.number {
                // 插到右边
                t.right = t.right.Insert(n)
        } else {
                // 插到左边
                t.left = t.left.Insert(n)
        }
        t.updateNode()
        return t.rotate()
}

// Remove 删除一个指定的元素（个数减一）
func (t *Tree) Remove(n int) (*Tree, error) {
        if t == nil {
                return nil, fmt.Errorf("Can not find number %d in AVL Tree", n)
        }
        if n == t.number {
                t.count--

                if t.count <= 0 {
                        if t.left != nil && t.right != nil {
                                minLeaf, tmp := t.right.removeLeft()
                                minLeaf.left = t.left
                                minLeaf.right = tmp
                                minLeaf.updateNode()

                                return minLeaf.rotate(), nil
                        } else if t.left != nil {
                                return t.left, nil
                        } else {
                                return t.right, nil
                        }
                }
                return t, nil
        } else if n > t.number {
                var err error
                t.right, err = t.right.Remove(n)
                t.right.updateNode()
                return t.rotate(), err
        } else {
                var err error
                t.left, err = t.left.Remove(n)
                t.left.updateNode()
                return t.rotate(), err
        }
}

func (t *Tree) removeLeft() (minLeaf *Tree, newRoot *Tree) {
        if t.left == nil {
                minLeaf = t
                newRoot = t.right
                return
        }
        minLeaf, t.left = t.left.removeLeft()
        t.updateNode()
        newRoot = t.rotate()
        return
}

// Min 获取平衡二叉树的最小值
func (t *Tree) Min() (int, error) {
        if t == nil {
                return 0, fmt.Errorf("Can not get the min number in an empty AVL Tree")
        }
        tmp := t
        for tmp.left != nil {
                tmp = tmp.left
        }
        return tmp.number, nil
}

// Max 获取平衡二叉树的最大值
func (t *Tree) Max() (int, error) {
        if t == nil {
                return 0, fmt.Errorf("Can not get the max number in an empty AVL Tree")
        }
        tmp := t
        for tmp.right != nil {
                tmp = tmp.right
        }
        return tmp.number, nil
}

// Depth 获取平衡二叉树节点的深度
func (t *Tree) Depth() int {
        if t == nil {
                return 0
        }
        return t.depth
}

// updateDepth 更新子树深度（左右子树深度必须已更新）
func (t *Tree) updateDepth() {
        if t != nil {
                leftDepth := t.left.Depth()
                rightDepth := t.right.Depth()

                if leftDepth < rightDepth {
                        t.depth = rightDepth + 1
                } else {
                        t.depth = leftDepth + 1
                }
        }
}

func (t *Tree) updateNode() {
        if t != nil {
                t.updateDepth()
        }
}

// Rotate 尝试旋转树（如果需要）
func (t *Tree) rotate() *Tree {
        if t == nil {
                return t
        }
        leftDepth := t.left.Depth()
        rightDepth := t.right.Depth()

        if leftDepth-rightDepth > 1 {
                // 左子树深度大于右子树深度
                // 右旋
                return t.rotateRight()
        } else if rightDepth-leftDepth > 1 {
                // 左子树深度小于右子树深度
                // 左旋
                return t.rotateLeft()
        }

        return t
}

// RotateRight 平衡二叉树右旋
func (t *Tree) rotateRight() *Tree {
        a := t
        b := a.left
        c := b.right

        a.left = c
        b.right = a

        a.updateNode()
        b.updateNode()

        return b
}

// RotateLeft 平衡二叉树左旋
func (t *Tree) rotateLeft() *Tree {
        a := t
        b := a.right
        c := b.left

        a.right = c
        b.left = a

        a.updateNode()
        b.updateNode()

        return b
}

// Count 获取指定节点存储的数字个数
func (t *Tree) Count() int {
        if t == nil {
                return 0
        }
        return t.count
}

// Number 获取指定节点存储的数字内容
func (t *Tree) Number() int {
        if t == nil {
                return 0
        }
        return t.number
}

// Print 打印二叉树
func (t *Tree) Print() {
        if t != nil {
                fmt.Printf("%d [%d %d] %d\n", t.Number(), t.left.Number(), t.right.Number(), t.Count())
                t.left.Print()
                t.right.Print()
        }
}
