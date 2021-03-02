package main

func main() {

}
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	s := make([]int, 0)
	t(k, 1, n, &res, s)
	return res
}
func t(k, index, n int, res *[][]int, s []int) {
	if len(s) == k {
		*res = append(*res, s)
	}

	for i := index; i <= n; i++ {
		slice := make([]int, len(s))
		copy(slice, s)
		slice = append(slice, i)
		t(k, i+1, n, res, slice)
		slice = slice[:len(slice)-1]
	}
}
