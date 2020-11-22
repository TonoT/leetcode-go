package main

import "sort"

func main() {
	findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}})
}
func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	num := 1
	slice := points[0]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= slice[1] {
			slice[0] = points[i][0]
			if points[i][1] <= slice[1] {
				slice[1] = points[i][1]
			}
		} else {
			num++
			slice = points[i]
		}
	}
	return num
}
