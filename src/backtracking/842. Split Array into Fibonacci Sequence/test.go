package main

import "math"

func main() {
	splitIntoFibonacci("0123")
}
func splitIntoFibonacci(S string) (F []int) {
	n := len(S)
	var backtrack func(index, sum, prev int) bool
	backtrack = func(index, sum, prev int) bool {
		if index == n {
			return len(F) >= 3
		}
		cur := 0
		for i := index; i < n; i++ {
			if i > index && S[index] == '0' {
				break
			}
			cur = cur*10 + int(S[i]-'0')
			if cur > math.MaxInt32 {
				break
			}
			if len(F) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}
			F = append(F, cur)
			if backtrack(i+1, prev+cur, cur) {
				return true
			}
			F = F[:len(F)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return
}
