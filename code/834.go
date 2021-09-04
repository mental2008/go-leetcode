var dp []int
var dp2 []int
var num []int
var next [][]int

func dfs(u, fa int) {
    num[u] = 1
    for _, v := range next[u] {
        if v == fa {
            continue
        }
        dfs(v, u)
        num[u] += num[v]
        dp[u] += (dp[v] + num[v])
    }
}

func dfs2(u, fa, n int) {
    if u == fa {
        dp2[u] = dp[u]
    } else {
        dp2[u] = dp2[fa] + n - 2*num[u]
    }
    for _, v := range next[u] {
        if v == fa {
            continue
        }
        dfs2(v, u, n)
    }
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
    dp, dp2, num, next = make([]int, n), make([]int, n), make([]int, n), make([][]int, n)
    for i := 0; i < n - 1; i++ {
        u, v := edges[i][0], edges[i][1]
        next[u] = append(next[u], v)
        next[v] = append(next[v], u)
    }
    dfs(0, 0)
    dfs2(0, 0, n)
    return dp2
}
