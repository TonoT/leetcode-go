package main

import "sort"

func main() {
	leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
}
func leastInterval(tasks []byte, n int) int {
	if len(tasks) <= 1 {
		return len(tasks)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] < tasks[j]
	})
	max := 1
	same := 1
	cnt := 1
	for i := 1; i < len(tasks); i++ {

		if tasks[i] == tasks[i-1] {
			cnt++
		}
		if tasks[i] != tasks[i-1] || i == len(tasks)-1 {
			if cnt > max {
				max = cnt
				same = 1
			} else if cnt == max {
				same++
			}
			cnt = 1
		}
	}
	if (max-1)*n+same-1 >= len(tasks)-max {
		return (max-1)*(n+1) + same
	} else {
		return len(tasks)
	}
}
