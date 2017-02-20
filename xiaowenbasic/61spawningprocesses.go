// 61spawningprocesses

/*
	> date
	Mon Feb 20 17:30:26 CST 2017

	> grep hello
	hello grep

	> ls -a -l -h
	total 2.5M
	drwxr-xr-x 1 Administrator 197121    0 Feb 20 17:30 .
	drwxr-xr-x 1 Administrator 197121    0 Feb 10 09:48 ..
	drwxr-xr-x 1 Administrator 197121    0 Feb 20 17:09 .git
	-rw-r--r-- 1 Administrator 197121  290 Jan 25 12:59 .gitignore
	-rw-r--r-- 1 Administrator 197121   95 Jan 25 12:59 README.md
	-rwxr-xr-x 1 Administrator 197121 2.5M Feb 20 17:30 gobyexample.exe
	-rw-r--r-- 1 Administrator 197121 1.2K Feb 20 17:30 gobyexample.go
	drwxr-xr-x 1 Administrator 197121    0 Feb 20 16:45 tmp
	drwxr-xr-x 1 Administrator 197121    0 Jan 25 13:01 xiaowenbasic

*/

package xiaowenbasic

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func SpawningProcessesMain() {
	fmt.Println("Sapwning Processes Main")

	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
