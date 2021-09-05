func orderlyQueue(s string, k int) string {
    if k >= 2 {
        r := []rune(s)
        sort.Slice(r, func(i, j int) bool {
            return r[i] < r[j]
        })
        return string(r)
    } else {
        start, n := 0, len(s)
        for i := 1; i < n; i++ {
            for j := 0; j < n; j++ {
                if s[(start+j)%n] < s[(i+j)%n] {
                    break
                }
                if s[(start+j)%n] > s[(i+j)%n] {
                    start = i
                    break
                }
            }
        }
        ans := ""
        for i := 0; i < n; i++ {
            ans += string(s[(start+i)%n])
        }
        return ans
    }
}
