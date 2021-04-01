func clumsy(N int) int {
    switch N {
        case 1:
            return 1
        case 2:
            return 2
        case 3:
            return 6
        case 4, 5 :
            return 7
        default:
            switch N % 4 {
                case 0:
                    return N+1
                case 3:
                    return N-1
                default:
                    return N+2
            }
    }
}
