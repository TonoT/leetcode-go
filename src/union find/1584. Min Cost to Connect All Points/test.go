package main

import "sort"

func main() {
	minCostConnectPoints2([][]int{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	})
}
func minCostConnectPoints(points [][]int) int {
	parent := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		parent[i] = i
	}
	l := make(map[int][][]int)
	length := make([]int, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			index := abs(points[j][0]-points[i][0]) + abs(points[j][1]-points[i][1])
			length = append(length, index)
			l[index] = append(l[index], []int{j, i})
		}
	}
	sort.Ints(length)
	s := -1
	all := 0
	for i := 0; i < len(length); i++ {
		if length[i] != s {
			s = length[i]
			for _, v := range l[length[i]] {
				if find(parent, v[0]) != find(parent, v[1]) {
					all += s
					union(parent, v[0], v[1])
				}
			}
		}
	}
	return all
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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

//prim
func minCostConnectPoints2(points [][]int) int {
	type point struct {
		pot int
		dis int
	}
	distance := make([]point, 0)
	//记录总长
	all := 0
	for i := 1; i < len(points); i++ {
		distance = append(distance, point{
			pot: i,
			dis: abs(points[i][0]-points[0][0]) + abs(points[i][1]-points[0][1]),
		})
	}

	for len(distance) > 0 {
		sort.Slice(distance, func(i, j int) bool {
			return distance[i].dis < distance[j].dis
		})
		all += distance[0].dis
		next := distance[0].pot
		for i := 1; i < len(distance); i++ {
			dis := abs(points[distance[i].pot][0]-points[next][0]) + abs(points[distance[i].pot][1]-points[next][1])
			if dis < distance[i].dis {
				distance[i].dis = dis
			}
		}
		distance = distance[1:]
	}
	return all
}
