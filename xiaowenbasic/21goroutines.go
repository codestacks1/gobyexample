// 21goroutines
// go例程
package xiaowenbasic

import (
	"fmt"
)

func GoroutinesMain() {
	fmt.Println("Goroutine")

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)  //等待读取用户录入
	fmt.Println("done") //进程完成
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
