type Char struct {
    chr string
    val int
}

var chars = []Char {
    {"M", 1000},
    {"CM", 900},
    {"D", 500},
    {"CD", 400},
    {"C", 100},
    {"XC", 90},
    {"L", 50},
    {"XL", 40},
    {"X", 10},
    {"IX", 9},
    {"V", 5},
    {"IV", 4},
    {"I", 1},
}

func romanToInt(s string) int {
    res := 0
    for _, c := range chars {

        for len(s) > 0 {
            if len(c.chr) == 2 && len(s) >= 2 && c.chr == s[:2] {
                res += c.val
                s = s[2:]
            } else if len(c.chr) == 1 && c.chr == s[:1] {
                res += c.val
                s = s[1:]
            } else {
                break
            }
        }
    }
    return res
}
