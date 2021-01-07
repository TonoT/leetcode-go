package main

func main() {
	findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	})
}
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	m := make([]int, n)
	for i := 0; i < n; i++ {
		m[i] = i
	}
	for i, v := range isConnected {
		for j, k := range v {
			if k == 1 && i != j {
				union(m, i, j)
			}
		}
	}
	cnt := 0
	for i, v := range m {
		if i == v {
			cnt += 1
		}
	}
	return cnt
}
func union(m []int, n1, n2 int) {
	m[find(m, n1)] = find(m, n2)
}
func find(m []int, n int) int {
	for m[n] != n {
		n = m[n]
		m[n] = m[m[n]]
	}
	return n
}

func findCircleNum2(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return
}
