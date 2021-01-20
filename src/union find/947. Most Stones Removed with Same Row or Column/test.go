package main

import "fmt"

//[[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
func main() {
	fmt.Print(removeStones2([][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 2},
		{2, 1},
		{2, 2},
	}))
}

func removeStones(stones [][]int) int {
	n := len(stones)
	graph := make([][]int, n)
	for i, p := range stones {
		for j, q := range stones {
			if p[0] == q[0] || p[1] == q[1] {
				graph[i] = append(graph[i], j)
			}
		}
	}
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		for _, w := range graph[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}
	cnt := 0
	for i, v := range vis {
		if !v {
			cnt++
			dfs(i)
		}
	}
	return n - cnt
}
func removeStones2(stones [][]int) int {
	graph := make([]int, 20000)
	for i := 0; i < len(graph); i++ {
		graph[i] = i
	}
	for _, v := range stones {
		union(graph, v[0], v[1]+10000)
	}
	cnt := 0
	m := make(map[int]struct{})
	for _, v := range stones {
		if graph[v[1]+10000] == v[1]+10000 {
			if _, ok := m[graph[v[1]+10000]]; !ok {
				cnt++
				m[graph[v[1]+10000]] = struct{}{}
			}
		}
	}
	return len(stones) - cnt
}
func union(p []int, n1, n2 int) {
	p[find(p, n1)] = find(p, n2)

}
func find(p []int, n int) int {
	for p[n] != n {
		p[n] = p[p[n]]
		n = p[n]
	}
	return n
}
