// 41defer
/*
	Defer被用来确保一个函数调用在程序执行结束前执行。
	同样用来执行一些清理工作，defer用在像其他语言中的ensure和finally用到的地方
*/
package xiaowenbasic

import (
	"fmt"
	"os"
)

func DeferMain() {
	fmt.Println("Defer main")

	f := createFile("tmp/defer.owntemp")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
