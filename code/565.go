func arrayNesting(nums []int) int {
    n := len(nums)
    vis := make([]bool, n)
    ans := 0
    for i := 0; i < n; i++ {
        if vis[i] {
            continue
        }
        res := 0
        for j := i; !vis[j]; j = nums[j] {
            res++
            vis[j] = true
        }
        if ans < res {
            ans = res
        }
    }
    return ans
}
