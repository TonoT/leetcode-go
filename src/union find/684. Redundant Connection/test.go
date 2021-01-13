package main

func main() {
	findRedundantConnection([][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	})
}
func findRedundantConnection(edges [][]int) []int {
	m := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		m[i] = i
	}
	for _, v := range edges {
		if find(m, v[0]-1) == find(m, v[1]-1) {
			return v
		} else {
			union(m, v[0]-1, v[1]-1)
		}
	}
	return nil
}
func union(p []int, n1, n2 int) {
	p[find(p, n1)] = find(p, n2)

}
func find(p []int, n int) int {
	for n != p[n] {
		n = p[n]
		p[n] = p[p[n]]
	}
	return n
}
