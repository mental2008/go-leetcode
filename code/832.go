func flipAndInvertImage(A [][]int) [][]int {
    n, m := len(A), len(A[0])
    swap := func(a int, b int) (int, int) {
        return b, a
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m/2; j++ {
            A[i][j], A[i][m-1-j] = swap(A[i][j], A[i][m-1-j])
            A[i][j], A[i][m-1-j] = 1-A[i][j], 1-A[i][m-1-j]
        }
        if m % 2 == 1 {
            A[i][m/2] = 1-A[i][m/2]
        }
    }
    return A
}
