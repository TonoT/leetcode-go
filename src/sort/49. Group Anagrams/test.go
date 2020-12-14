package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, i := range strs {
		cnt := [26]int{}
		for _, v := range i {
			cnt[v-'a']++
		}
		m[cnt] = append(m[cnt], i)
	}
	res := make([][]string, 0, len(m))
	for _, i := range m {
		res = append(res, i)
	}
	return res
}
