// 58commandlinearguments
package xiaowenbasic

import (
	"fmt"
	"os"
)

func CommandLineArgumentsMain() {
	fmt.Println("Command Line Arguments main")

	argsWithProg := os.Args //注意：参数可能为空
	argsWithoutProg := os.Args[1:]
	//	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	//	fmt.Println(arg)
}
