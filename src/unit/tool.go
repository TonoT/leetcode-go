package unit

import (
	"fmt"
	"strings"
)

//func main() {
//	changetoslice("[[9,12],[1,10],[4,11],[8,12],[3,9],[6,9],[6,7]]")
//}

func changetoslice(s string) {
	ns := strings.ReplaceAll(s, "[", "{")
	ns = strings.ReplaceAll(ns, "]", "}")
	fmt.Print(ns)
}
