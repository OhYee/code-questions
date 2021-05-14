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

func intToRoman(num int) string {
    res := ""

    for _, c := range chars {
        for num >= c.val {
            num -= c.val
            res += c.chr
        }
    }

    return res
}
