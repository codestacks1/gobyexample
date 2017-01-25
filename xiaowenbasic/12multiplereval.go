// 12MultipleRetVal
package xiaowenbasic

import (
	"fmt"
)

func MultipleRetValmain() {
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

func vals() (int, int) {
	return 2, 7
}
