// Panic
/*
	panic 意味着有些出乎意料的错误发生。
	通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误
*/
package xiaowenbasic

import (
	"fmt"
	"os"
)

func PanicMain() {
	fmt.Println("Panic main")

	panic("a problem") // 该方法会启动一个异常

	_, err := os.Create("tmp/temp.ownmd")
	if err != nil {
		panic(err)
	}
}
