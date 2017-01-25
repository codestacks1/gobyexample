// 25channelDirections
package xiaowenbasic

import (
	"fmt"
)

func ChannelDirectionsMain() {
	fmt.Println("Channel Direction")

	pings := make(chan string, 1)
	pongs := make(chan string, 2)

	ping(pings, "Passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
