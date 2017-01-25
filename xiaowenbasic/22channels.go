// 22channels
package xiaowenbasic

import (
	"fmt"
)

func ChannelMain() {
	fmt.Println("Channel")

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
