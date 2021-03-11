type OpTree interface {
	Calc() int
	Debug() string
}

type Number struct {
	num int
}

func NewNumber(num int) *Number {
	return &Number{
		num: num,
	}
}

func (t *Number) Calc() int {
	if t == nil {
		return 0
	}
	return t.num
}

func (t *Number) Debug() string {
	if t == nil {
		return ""
	}
	return fmt.Sprint(t.num)
}

type Operator struct {
	L      OpTree
	R      OpTree
	Op     byte
	tempOp byte
}

func (t *Operator) Calc() (res int) {
	if t == nil {
		return 0
	}

	var l, r int
	if t.L != nil {
		l = t.L.Calc()
	}
	if t.R != nil {
		r = t.R.Calc()
	}

	switch t.Op {
	case '+':
		res = l + r
	case '-':
		res = l - r
	case '*':
		res = l * r
	case '/':
		res = l / r
	default:
		res = l
	}
	// fmt.Printf("%d %c %d = %d\n", l, t.Op, r, res)
	return res
}

func (t *Operator) Debug() string {
	if t == nil {
		return ""
	}
	l := ""
	r := ""
	if t.L != nil {
		l = t.L.Debug()
	}
	if t.R != nil {
		r = t.R.Debug()
	}
	return fmt.Sprintf("(%s) %c (%s)", l, t.Op, r)
}

func (t *Operator) Insert(nt OpTree, op byte) *Operator {
	if t == nil {
		t = new(Operator)
	}

	// ???
	fmt.Print()

	if t.L == nil {
		t.L = nt
		t.Op = op
		return t
	} else if t.R == nil {
		t.R = nt
		t.tempOp = op
		return t
	}
	// 两侧都已满
	root := &Operator{
		L:      t,
		R:      nt,
		Op:     t.tempOp,
		tempOp: op,
	}
	return root
}

func Parse(s string, start, n int) (root *Operator, idx int) {
	root = new(Operator)
	var temp *Operator = nil
	num := 0
	hasNum := false

	getLeft := func() (left OpTree) {
		if hasNum {
			left = NewNumber(num)
			num = 0
			hasNum = false
		}
		if temp != nil {
			if left != nil {
				temp = temp.Insert(left, '[')
			}
			left = temp
			temp = nil
		}
		return
	}

	defer func() {
		root = root.Insert(
			getLeft(),
			']',
		)
		// if root.R == nil {
		// 	root = root.L.(*Operator)
		// }
	}()

	for i := start; i < n; i++ {
		switch s[i] {
		case '(':
			temp, i = Parse(s, i+1, n)
		case ')':
			return root, i
		case '+', '-':
			root = root.Insert(
				getLeft(),
				s[i],
			)
		case '*', '/':
			temp = temp.Insert(
				getLeft(),
				s[i],
			)
		case ' ':
			continue
		default:
			// 数字
			hasNum = true
			num = num*10 + int(s[i]-'0')
		}
	}
	return root, n
}

func calculate(s string) int {
	n := len(s)
	root, _ := Parse(s, 0, n)
	// fmt.Println(root.Debug())
	return root.Calc()
}
