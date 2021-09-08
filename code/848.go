func shiftingLetters(s string, shifts []int) string {
    n := len(s)
    ss := []byte(s)
    for i := n-1; i >= 0; i-- {
        if i < n-1 {
            shifts[i] += shifts[i+1]
        }
        ss[i] = byte('a' + (int(ss[i] - 'a') + shifts[i]) % 26)
    }
    return string(ss)
}
