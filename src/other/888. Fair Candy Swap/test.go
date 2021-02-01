package main

func main() {

}
func fairCandySwap(A []int, B []int) []int {
	sum := func(s []int) int {
		total := 0
		for _, v := range s {
			total += v
		}
		return total
	}
	i := sum(A)
	j := sum(B)
	m := make(map[int]bool)
	for _, v := range B {
		m[v] = true
	}
	avg := (i + j) / 2
	for _, v := range A {
		if _, ok := m[v+(avg-i)]; ok {
			return []int{v, v + (avg - i)}
		}
	}
	return nil
}
