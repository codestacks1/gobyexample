// 14Closures
// 闭包
package xiaowenbasic

import (
	"fmt"
)

func ClosuresMain() {
	nextInt := intSeq()

	fmt.Println(nextInt)
	fmt.Println(nextInt)
	fmt.Println(nextInt)

	newInts := intSeq()
	fmt.Println(newInts)   //output 0x45d3f0
	fmt.Println(newInts()) //output 1
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
