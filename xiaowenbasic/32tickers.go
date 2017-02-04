// 32tickers
package xiaowenbasic

import (
	"fmt"
	"time"
)

func TickersMain() {
	fmt.Println("Hello World!")

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
