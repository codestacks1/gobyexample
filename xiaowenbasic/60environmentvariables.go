// 60environmentvariables
/*
	该代码内容，获取环境变量
*/
package xiaowenbasic

import (
	"fmt"
	"os"
	"strings"
)

func EnvironmentVariablesMain() {
	fmt.Println("Environment Variables Main")

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
