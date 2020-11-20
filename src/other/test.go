package main

import (
	"fmt"
	"strings"
)

func backspaceCompare(S string, T string) bool {
	return strings.EqualFold(cal(S),cal(T))
}
func cal(s string)string{
	b:=[]byte("#")[0]
	sa:=[]byte(s)
	nb:=make([]byte,0)
	for i, _ := range sa {
		if sa[i]!=b{
			nb=append(nb,s[i])
		}else{
			if len(nb)>0{
				nb=nb[:len(nb)-1]
			}
		}
	}
	return string(nb)
}
func main() {
fmt.Println(cal("##sa#a#ss"))
}
