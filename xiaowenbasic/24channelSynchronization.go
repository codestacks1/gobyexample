// 24channelSynchronization
package xiaowenbasic

import (
	"fmt"
	"time"
)

func ChannelSynchronizationMain() {
	fmt.Println("Channel Synchronization")

	done := make(chan bool, 10)
	go worker(done)

	<-done
}

func worker(done chan bool) {
	fmt.Println("working...", time.Now())
	time.Sleep(time.Second)
	fmt.Println("done", "\t", time.Now())

	done <- true
}
