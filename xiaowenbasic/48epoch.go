// 48epoch
/*
	时期，纪元
	output:
			Epoch main
			2017-02-06 13:12:56.9116001 +0800 CST
			1486357976
			1486357976911
			1486357976911600100
			2017-02-06 13:12:56 +0800 CST
			2017-02-06 13:12:56.9116001 +0800 CST
*/

package xiaowenbasic

import (
	"fmt"
	"time"
)

func EpochMain() {
	fmt.Println("Epoch main")

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
