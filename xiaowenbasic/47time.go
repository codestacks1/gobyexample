// 47time
/*
	output:
			Time main
			2017-02-06 13:08:29.5296001 +0800 CST
			2017-11-17 20:34:58.651387237 +0000 UTC
			2017
			November
			17
			20
			34
			58
			651387237
			UTC
			Friday
			false
			true
			false
			-6831h26m29.121787137s
			-6831.441422718649
			-409886.48536311893
			-2.459318912178714e+07
			-24593189121787137
			2017-02-06 05:08:29.5296001 +0000 UTC
			2018-08-29 12:01:27.773174374 +0000 UTC
*/
package xiaowenbasic

import (
	"fmt"
	"time"
)

func TimeMain() {
	fmt.Println("Time main")

	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2017, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
