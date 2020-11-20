package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(reconstructQueue([][]int{
		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1}, {5, 0}, {6, 1}, {5, 2},
	}))
}

//Input:
//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
//Output:
//[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
//for _, person := range people {
//spaces := person[1] + 1
//for i := range ans {
//if ans[i] == nil {
//spaces--
//if spaces == 0 {
//ans[i] = person
//break
//}
//}
//}
//}
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || people[i][0] == people[j][0] && people[i][1] > people[j][1]
	})
	ans := make([][]int, len(people))
	for _, person := range people {
		count := person[1] + 1
		for i, _ := range ans {
			if ans[i] == nil {
				count--
				if count == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}
