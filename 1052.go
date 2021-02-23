func maxSatisfied(customers []int, grumpy []int, X int) int {
    len := len(customers)
    f, g := make([]int, len+1), make([]int, len+1)
    f[0], g[0] = 0, 0
    for i := 0; i < len; i++ {
        f[i+1] = f[i] + customers[i]
        g[i+1] = g[i] + customers[i] *  (1 - grumpy[i])
    }
    ans := 0
    for i := 0; i <= len-X; i++ {
        res := f[i+X] - f[i] + g[len] - g[i+X] + g[i]
        if ans < res {
            ans = res
        }
    }
    return ans
}
