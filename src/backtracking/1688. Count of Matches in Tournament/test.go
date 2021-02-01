package main

func main() {
	numberOfMatches(7)

}
func numberOfMatches(n int) int {
	y := n
	sum := 0
	for y > 1 {
		if y%2 == 0 {
			sum += y >> 1
			y = y >> 1
		} else {
			sum += y >> 1
			y = y>>1 + 1
		}
	}
	return sum
}
