func slowestKey(releaseTimes []int, keysPressed string) byte {
    n := len(releaseTimes)
    l, c := releaseTimes[0], keysPressed[0]
    for i := 1; i < n; i++ {
        if (releaseTimes[i] - releaseTimes[i-1] > l) || (releaseTimes[i] - releaseTimes[i-1] == l && keysPressed[i] > c) {
            c = keysPressed[i]
            l = releaseTimes[i] - releaseTimes[i-1]
        }
    }
    return c
}
