// 63signals
package xiaowenbasic

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func SignalsMain() {
	fmt.Println("Signals Main")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//go routines 好像类似C#Action委托 new Action=>()=>{}()
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
