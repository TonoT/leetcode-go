package main

func main() {

}
func fib(n int) int {
	a := 0
	b := 1
	m := 0
	for i := 2; i <= n; i++ {
		m = a + b
		a = b
		b = m
	}
	if n < 2 {
		return n
	} else {
		return m
	}
}
