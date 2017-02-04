// 37statefulgoroutines
package xiaowenbasic

import (
	"fmt"
)

func StatefulGoroutinesMain() {
	fmt.Println("Stateful Goroutines main")

	messages := make(chan string) // channel

	go func() { // go routines
		messages <- "ping" // go channel
	}()

	msg := <-messages // go channel
	fmt.Println(msg)
}
