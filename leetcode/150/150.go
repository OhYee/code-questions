func evalRPN(tokens []string) int {
    stack := make([]int, len(tokens))
    top := 0

    pop := func() int {
        res := stack[top-1]
        top--
        return res
    }
    push := func(num int) {
        stack[top] = num
        top++
    }
    get2Numbers := func() (int, int) {
        b := pop()
        a := pop()
        return a, b
    }

    for _, token := range tokens {
        switch token {
            case "+":
                a, b := get2Numbers()
                push(a+b)
            case "-":
                a, b := get2Numbers()
                push(a-b)
            case "*":
                a, b := get2Numbers()
                push(a*b)
            case "/":
                a, b := get2Numbers()
                push(a/b)
            default:
                a, _ := strconv.ParseInt(token, 10, 64)
                push(int(a))
        }
    }
    return stack[0]
}
