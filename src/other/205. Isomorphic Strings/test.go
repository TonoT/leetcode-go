package main

func main() {
	isIsomorphic("foo", "bar")
}
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[int32]int32)
	for i, v := range t {
		if n, ok := m[v]; ok {
			if n != int32(s[i]) {
				return false
			}
		} else {
			m[v] = int32(s[i])
		}
	}
	m2 := make(map[int32]int32)
	for i, v := range s {
		if n, ok := m2[v]; ok {
			if n != int32(t[i]) {
				return false
			}
		} else {
			m2[v] = int32(t[i])
		}
	}
	return true
}
