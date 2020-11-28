package main

import "fmt"

func reversePairs(nums []int) int {
	t := NewTree()
	res := 0
	for _, n := range nums {
		q := t.Query(n*2 + 1)
		res += q
		t = t.Insert(n)
		// t.Print()
		// fmt.Printf("----[%d %d]----\n\n", n, q)
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
	sum    int
}

// NewTree 生成一个平衡二叉树
func NewTree() *Tree {
	return nil
}

func NewTreeNode(number int) *Tree {
	return &Tree{
		number: number,
		count:  1,
		left:   nil,
		right:  nil,
		depth:  1,
		sum:    1,
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

// Query 检索所有大于该值的数个数
func (t *Tree) Query(n int) int {
	if t == nil {
		return 0
	}
	if n == t.number {
		return t.Count() + t.right.Query(n)
	} else if n > t.number {
		return t.right.Query(n)
	} else {
		return t.Count() + t.left.Query(n) + t.right.Sum()
	}
}

func (t *Tree) Depth() int {
	if t == nil {
		return 0
	}
	return t.depth
}

func (t *Tree) Count() int {
	if t == nil {
		return 0
	}
	return t.count
}

func (t *Tree) Sum() int {
	if t == nil {
		return 0
	}
	return t.sum
}

func (t *Tree) Number() int {
	if t == nil {
		return 0
	}
	return t.number
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

func (t *Tree) updateSum() {
	if t != nil {
		t.sum = t.Count() + t.left.Sum() + t.right.Sum()
	}
}

func (t *Tree) updateNode() {
	if t != nil {
		t.updateDepth()
		t.updateSum()
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

func (t *Tree) Print() {
	if t != nil {
		fmt.Printf("%d [%d %d] %d %d\n", t.Number(), t.left.Number(), t.right.Number(), t.Count(), t.Sum())
		t.left.Print()
		t.right.Print()
	}
}
