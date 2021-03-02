package main

import "fmt"

func main() {
	possibilities := numTilePossibilities("AAB")
	fmt.Println(possibilities)

}

func numTilePossibilities(tiles string) int {
	s := make([]int, 26)
	for i, _ := range tiles {
		s[tiles[i]-'A']++
	}
	return dfs(s)
}
func dfs(s []int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			continue
		}
		res += 1
		s[i]--
		res += dfs(s)
		s[i]++
	}
	return res
}
