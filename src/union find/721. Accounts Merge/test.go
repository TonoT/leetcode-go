package main

func main() {

}
func accountsMerge(accounts [][]string) (res [][]string) {
	name := make(map[string]string)
	index := make(map[string]int)
	for _, v := range accounts {
		n := v[0]
		for _, l := range v[1:] {
			name[l] = n
			if _, ok := index[l]; !ok {
				index[l] = len(index)
			}

		}
	}
	parent := make([]int, len(index))
	for i := 0; i < len(index); i++ {
		parent[i] = i
	}
	for _, account := range accounts {
		first := account[1]
		for _, n := range account[2:] {
			union(parent, index[n], index[first])
		}
	}

	sm := make(map[int][]string)
	for n, i := range index {
		num := find(parent, i)
		sm[num] = append(sm[num], n)
	}
	for _, v := range sm {
		s := append([]string{name[v[0]]}, v...)
		res = append(res, s)
	}
	return
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
