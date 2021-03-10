type OpTree struct {
    L       *OpTree
    R       *OpTree
    Op      string
    IsNum   bool
    Num     int
}

func NewNum(num int) *OpTree {
    return &OpTree {
        IsNum: true,
        Num: num,
    }
}

func (t *OpTree) Calc() int {
    if t == nil {
        return 0
    }
    if t.IsNum {
        return t.Num
    }

    l := t.L.Calc()
    r := t.R.Calc()
    switch t.Op {
        case "+":
            return l + r
        case "-":
            return l - r
    }
    return l
}

func (t *OpTree) Insert(nt *OpTree) *OpTree {
    if t.L == nil {
        t.L = nt
        return t
    } 
    t.R = nt
    root := new(OpTree)
    root.L = t
    return root
}

func (t *OpTree) Debug() string {
    if t == nil {
        return ""
    }
    if t.IsNum {
        return fmt.Sprintf("%d", t.Num)
    }
    return fmt.Sprintf("(%s) %s (%s)", t.L.Debug(), t.Op, t.R.Debug())
}

func Parse(s string, start, n int ) (root *OpTree, idx int) {
    root = new(OpTree)
    num := 0
    defer func () {
        // 结束前需要把最后一个数字也加进去
        root = root.Insert(NewNum(num))
    }()
    for i := start; i < n; i++ {
        switch s[i] {
            case '(':
                c, j := Parse(s, i+1, n)
                root = root.Insert(c)
                // fmt.Println(s[i:j+1], c)
                i = j
            case ')':
                return root, i
            case '+':
                root = root.Insert(
                    NewNum(num),
                )
                root.Op = "+"
                num = 0
            case '-':
                root = root.Insert(
                    NewNum(num),
                )
                root.Op = "-"
                num = 0
            case ' ':
                continue
            default:
                // 数字
                num = num * 10 + int(s[i] - '0')
        }
    }
    return root, n
}

func calculate(s string) int {
    // 先消去所有的空格
    n := len(s)
    root, _ := Parse(s, 0, n)
    // fmt.Println(root.Debug())
    return root.Calc()
}
