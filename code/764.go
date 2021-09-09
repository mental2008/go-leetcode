func orderOfLargestPlusSign(n int, mines [][]int) int {
    a, up, down, left, right := make([][]int, n), make([][]int, n), make([][]int, n), make([][]int, n), make([][]int, n)
    for i := 0; i < n; i++ {
        a[i] = make([]int, n)
        up[i] = make([]int, n)
        down[i] = make([]int, n)
        left[i] = make([]int, n)
        right[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            a[i][j] = 1
        }
    }
    for i, _ := range(mines) {
        x, y := mines[i][0], mines[i][1]
        a[x][y] = 0
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if a[i][j] == 0 {
                right[i][j] = 0
                continue
            }
            if j == 0 {
                right[i][j] = 1
            } else {
                right[i][j] = right[i][j-1] + 1
            }
        }
        for j := n-1; j >= 0; j-- {
            if a[i][j] == 0 {
                left[i][j] = 0
                continue
            }
            if j == n-1 {
                left[i][j] = 1
            } else {
                left[i][j] = left[i][j+1] + 1
            }
        }
    }
    for j := 0; j < n; j++ {
        for i := 0; i < n; i++ {
            if a[i][j] == 0 {
                up[i][j] = 0
                continue
            }
            if i == 0 {
                up[i][j] = 1
            } else {
                up[i][j] = up[i-1][j] + 1
            }
        }
        for i := n-1; i >= 0; i-- {
            if a[i][j] == 0 {
                down[i][j] = 0
                continue
            }
            if i == n-1 {
                down[i][j] = 1
            } else {
                down[i][j] = down[i+1][j] + 1
            }
        }
    }
    ans := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            res := up[i][j]
            if res > down[i][j] {
                res = down[i][j]
            }
            if res > left[i][j] {
                res = left[i][j]
            }
            if res > right[i][j] {
                res = right[i][j]
            }
            if ans < res {
                ans = res
            }
        }
    }
    return ans
}
