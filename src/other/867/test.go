package main

func main() {

}
func transpose(matrix [][]int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(matrix[0]); i++ {
		s := make([]int, 0)
		for j := 0; j < len(matrix); j++ {
			s = append(s, matrix[j][i])
		}
		res = append(res, s)
	}
	return res
}
