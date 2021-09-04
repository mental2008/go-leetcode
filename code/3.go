func lengthOfLongestSubstring(s string) int {
    mp := make(map[rune]int)
    start, ans := 0, 0
    for i, val := range s {
        pos := mp[val]
        if pos != 0 {
            for j := start+1; j <= pos; j++ {
                mp[rune(s[j-1])] = 0
            }
            start = pos
        }
        mp[val] = i+1
        if i+1-start > ans {
            ans = i+1-start
        }
    }
    return ans
}
