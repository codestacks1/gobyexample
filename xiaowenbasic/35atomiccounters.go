// 35atomiccounters
// 原子计数器
/*
	Go中最重要的状态管理是通过通道间的沟通来完成的，我们在workerpools的例子中碰到过，
	但还是有一些其他的方法来管理状态的，这里我们将看看如何使用sync/atomic包在多个Go协程中进行院子计数。
*/
package xiaowenbasic

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func AtomicCountersMain() {
	fmt.Println("Atomic Counters main")

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			fmt.Println(">> ", i)
			atomic.AddUint64(&ops, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
