// 64exit
package xiaowenbasic

import (
	"fmt"
	"os"
)

func ExitMain() {
	fmt.Println("Exit Main")

	defer fmt.Println("!")

	os.Exit(31) //杀死当前进程
}
