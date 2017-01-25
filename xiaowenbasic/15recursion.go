// 15Recursion
// 递归
package xiaowenbasic

import (
	"fmt"
)

func RecursionMain() {
	fmt.Println(fact(7))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return n * fact(n-1)
}
