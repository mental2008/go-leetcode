func isLower(a, b, c []int) bool {
    return (c[1]-a[1])*(b[0]-a[0]) < (b[1]-a[1])*(c[0]-a[0])
}

func isUpper(a, b, c []int) bool {
    return (c[1]-a[1])*(b[0]-a[0]) > (b[1]-a[1])*(c[0]-a[0])
}

func outerTrees(trees [][]int) [][]int {
    n := len(trees)
    if n == 1 {
        return trees
    }

    sort.Slice(trees, func(i, j int) bool {
        if trees[i][0] == trees[j][0] {
            return trees[i][1] < trees[j][1]
        }
        return trees[i][0] < trees[j][0]
    })

    ans, vis := make([][]int, 0), make([]bool, n)
    for i, _ := range vis {
        vis[i] = false
    }

    lres := []int{0, 1}
    m := 2
    for i := 2; i < n; i++ {
        for m>=2 && isLower(trees[lres[m-2]], trees[lres[m-1]], trees[i]) {
            m--
        }
        lres = append(lres[:m], i)
        m++
    }
    k := 0
    for i := 0; i < m; i++ {
        vis[lres[i]] = true
        ans = append(ans, make([]int, 2))
        ans[k] = trees[lres[i]]
        k++
    }

    ures := []int{0, 1}
    m = 2
    for i := 2; i < n; i++ {
        for m>=2 && isUpper(trees[ures[m-2]], trees[ures[m-1]], trees[i]) {
            m--
        }
        ures = append(ures[:m], i)
        m++
    }
    for i := 0; i < m; i++ {
        if vis[ures[i]] {
           continue
        }
        vis[ures[i]] = true
        ans = append(ans, make([]int, 2))
        ans[k] = trees[ures[i]]
        k++
    }
    return ans
}
