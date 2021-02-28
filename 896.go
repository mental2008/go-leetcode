func isMonotonic(A []int) bool {
    t := 0
    for i := 1; i < len(A); i++ {
        if A[i-1] < A[i] {
            if t == -1 {
                return false
            }
            t = 1
        } else if A[i-1] > A[i] {
            if t == 1 {
                return false
            }
            t = -1
        }
    }
    return true
}
