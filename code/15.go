func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    ans := make([][]int, 0)
    for i := 0; i < n; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        k := n-1
        for j := i+1; j < n; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            for j < k && nums[i] + nums[j] + nums[k] > 0 {
                k--
            }
            if j == k {
                break
            }
            if nums[i] + nums[j] + nums[k] == 0 {
                ans = append(ans, []int{nums[i], nums[j], nums[k]})
            }
        }
    }
    return ans
}
