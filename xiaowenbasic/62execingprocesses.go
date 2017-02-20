// execingprocesses

/*
	注意：
		一下操作，不被Windows系统支持
*/

package xiaowenbasic

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func ExecingProcessesMain() {
	fmt.Println("Execing Processes Main")

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
