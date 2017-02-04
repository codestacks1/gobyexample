// 44stringformatting
package xiaowenbasic

import (
	"fmt"
	"os"
)

func StringFormattingMain() {
	fmt.Println("String Formatting main")

	p := point{x: 1, y: 2}
	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)
}

type point struct {
	x, y int
}
