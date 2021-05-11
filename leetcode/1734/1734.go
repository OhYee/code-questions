func decode(encoded []int) []int {
    n := len(encoded) + 1
    perm := make([]int, n)

    abcdefg := 0
    bcdefg := 0
    for i := 0; i < n; i++ {
        abcdefg ^= i + 1
        if i % 2 == 1 {
            bcdefg ^= encoded[i]
        }
    }

    perm[0] = abcdefg ^ bcdefg
    for i := 1; i < n; i++ {
        perm[i] = encoded[i-1] ^ perm[i-1]
    }

    r
