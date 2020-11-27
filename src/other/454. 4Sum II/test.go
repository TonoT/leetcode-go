package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	ans := 0
	sum := make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			sum[A[i]+B[j]]++
		}
	}
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			ans += sum[-(C[i] + D[j])]
		}
	}
	return ans
}
