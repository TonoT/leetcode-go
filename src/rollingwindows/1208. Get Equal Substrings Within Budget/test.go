package main

func main() {

}
func equalSubstring(s string, t string, maxCost int) int {
	dis := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if t[i] > s[i] {
			dis[i] = int(t[i] - s[i])
		} else {
			dis[i] = int(s[i] - t[i])
		}

	}
	r, max := 0, 0
	for _, v := range dis {
		max += v
		if max > maxCost {
			max -= dis[r]
			r++
		}

	}
	return len(s) - r
}
