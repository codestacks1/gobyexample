// 49timeformatting/parsing

/*
	output:
			Time Formatting / Parsing main
			2017-02-06T13:26:42+08:00
			2012-11-01 22:08:41 +0000 +0000
			1:26PM
			Mon Feb  6 13:26:42 2017
			2017-02-06T13:26:42.1446+08:00
			0000-01-01 20:41:00 +0000 UTC
			2017-02-06T13:26:42-00:00
			parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
*/
package xiaowenbasic

import (
	"fmt"
	"time"
)

func TimeFormattingParsingMain() {
	fmt.Println("Time Formatting / Parsing main")

	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
