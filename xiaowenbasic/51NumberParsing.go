// 51NumberParsing
/*
	output:
			Number Parsing main
			1.234
			123
			456
			0
			135
			strconv.ParseInt: parsing "wat": invalid syntax
*/

package xiaowenbasic

import (
	"fmt"
	"strconv"
)

func NumberParsingMain() {
	fmt.Println("Number Parsing main")

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("-789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
