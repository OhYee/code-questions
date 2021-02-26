import (
    "math/bits"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
    m := make(map[int]int)
    for _, word := range words{
        num := 0
        for _, c := range(word) {
            num |= 1<<(c - 'a')
        }
        if bits.OnesCount(uint(num)) <= 7 {
            m[num]++
        }
    }
    res := make([]int, len(puzzles))
    for idx, puzzle := range(puzzles) {
        count := 0
        for i:=0; i<(1<<6); i++ {
            temp := 1 << (puzzle[0] - 'a')
            for j:=0; j<6; j++ {
                temp |= ((i>>j)&1)<<(puzzle[j+1] - 'a')
            }
            count += m[temp]
        }
        res[idx] = count
    }
    return res
}
