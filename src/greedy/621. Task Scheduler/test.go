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

func leastInterval2(tasks []byte, n int) int {
	if len(tasks) <= 1 {
		return len(tasks)
	}
	max := 0
	same := 1
	m := map[byte]int{}
	for _, v := range tasks {
		m[v]++
	}
	for _, v := range m {
		if v > max {
			max = v
			same = 1
		} else if v == max {
			same++
		}
	}
	if (max-1)*n+same-1 >= len(tasks)-max {
		return (max-1)*(n+1) + same
	} else {
		return len(tasks)
	}
}
func leastInterval3(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	bucket := make([]int, 26)
	for i := range tasks {
		bucket[tasks[i]-'A']++
	}
	sort.Ints(bucket)
	max := bucket[len(bucket)-1]
	j := 1
	res := (max-1)*n + max
	for i := len(bucket) - 2; i >= 0 && bucket[i] > 0; i-- {
		temp := (bucket[i]-1)*n + bucket[i] + j
		if temp > res {
			res = temp
		} else {
			break
		}
		j++
	}
	if res < len(tasks) {
		return len(tasks)
	}
	return res
}
func leastInterval4(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	bucket := make([]int, 26)
	for i := range tasks {
		bucket[tasks[i]-'A']++
	}
	sort.Ints(bucket)
	same := 1
	max := bucket[len(bucket)-1]
	for i := len(bucket) - 2; i > 0; i-- {
		if bucket[i] == max {
			same++
		}
	}
	if (max-1)*n+same-1 >= len(tasks)-max {
		return (max-1)*(n+1) + same
	} else {
		return len(tasks)
	}
}
