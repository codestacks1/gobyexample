// 38sorting
package xiaowenbasic

import (
	"fmt"
	"sort"
)

func SortingMain() {
	fmt.Println("Go Sorting main")

	strs := []string{"c", "a", "b", "f", "e"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{3, 1, 6, 10, 2}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
