func isMonotonic(A []int) bool {
    status := 0
    for i, num := range A {
        if i == 0 {
            continue
        }
        if status == 0 {
            if num > A[i-1] {
                status = 1;
            } else if num < A[i-1] {
                status = -1;
            }
        } else if status == 1 && num < A[i-1] {
            return false
        } else if status == -1 && num > A[i-1] {
            return false
        }
    }
    return true
}
