type NumMatrix struct {
    sum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    n := len(matrix)
    if n == 0 {
        return NumMatrix{}
    }
    m := len(matrix[0])
    sum := make([][]int, n+1)
    sum[0] = make([]int, m+1)
    for i := 1; i <= n; i++ {
        sum[i] = make([]int, m+1)
        for j := 1; j <= m; j++ {
            sum[i][j] = sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1] + matrix[i-1][j-1]
        }
    }
    return NumMatrix{sum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.sum[row2+1][col2+1] - this.sum[row2+1][col1] - this.sum[row1][col2+1] + this.sum[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
