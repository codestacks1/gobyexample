// 36mutexes
/*
	互斥锁，安全的在Go协程间访问数据
*/
package xiaowenbasic

import (
	"fmt"
)

func MutexesMain() {
	fmt.Println("Mutexes main")

	fu("direct")

	go fu("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done.")
}

func fu(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
