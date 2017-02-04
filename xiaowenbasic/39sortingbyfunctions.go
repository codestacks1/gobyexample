// 39sortingbyfunctions
package xiaowenbasic

import (
	"fmt"
	"sort"
)

func SortingByFunctionsMain() {
	fmt.Println("Sorting Functions main")

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

type ByLength []string

func (s ByLength) Len() int {
	//	fmt.Println(len(s))
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	//	fmt.Println(s[j] + "..." + s[i])
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
