func transpose(matrix [][]int) [][]int {
    n, m := len(matrix), len(matrix[0])
    trans := make([][]int, m)
    for j := 0; j < m; j++ {
        trans[j] = make([]int, n)
        for i := 0; i < n; i++ {
            trans[j][i] = matrix[i][j]
        }
    }
    return trans
}
