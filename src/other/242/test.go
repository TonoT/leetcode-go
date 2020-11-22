package main

import "sort"

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sByte := []byte(s)
	tByte := []byte(t)
	sort.Slice(sByte, func(i, j int) bool {
		return sByte[i] < sByte[j]
	})
	sort.Slice(tByte, func(i, j int) bool {
		return tByte[i] < tByte[j]
	})
	return string(sByte) == string(tByte)
}
