// 34ratelimiting
// 速率限制
// 速率限制，是一个重要的控制服务资源利用和质量的途径，Go通过Go协程、通道和打点器优美的支持了速率限制
package xiaowenbasic

import (
	"fmt"
	"time"
)

func RateLimitingMain() {
	fmt.Println("Rate Limiting main")

	requests := make(chan int, 5) //声明int类型的channel,容量为5,
	for i := 1; i <= 5; i++ {
		requests <- i //为channel赋值
	}

	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
