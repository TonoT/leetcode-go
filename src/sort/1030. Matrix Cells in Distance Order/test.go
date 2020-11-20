package main

import (
	"math"
	"sort"
)

func main() {

}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	ans := make([][]int, 0)
	if R == 0 || C == 0 {
		return nil
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			ans = append(ans, []int{i, j})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		f := math.Abs(float64(ans[i][0]-r0)) + math.Abs(float64(ans[i][1]-c0))
		s := math.Abs(float64(ans[j][0]-r0)) + math.Abs(float64(ans[j][1]-c0))
		return f < s
	})
	return ans
}
