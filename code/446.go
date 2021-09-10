func numberOfArithmeticSlices(nums []int) int {
    n := len(nums)
    dp := make([][]int, n)
    mp := make([]map[int]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        mp[i] = make(map[int]int, 0)
    }
    ans := 0
    for i := n-1; i >= 0; i-- {
        for j := i+1; j < n; j++ {
            dp[i][j] = mp[j][nums[j]-nums[i]]
            ans += dp[i][j]
        }
        for j := i+1; j < n; j++ {
            mp[i][nums[j]-nums[i]]+=(dp[i][j]+1)
        }
    }
    return ans
}
