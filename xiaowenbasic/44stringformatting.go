// 44stringformatting
package xiaowenbasic

import (
	"fmt"
	"os"
)

func StringFormattingMain() {
	fmt.Println("String Formatting main")

	p := point{x: 1, y: 2}
	fmt.Printf("%v\n", p)       //value
	fmt.Printf("%+v\n", p)      //key + value
	fmt.Printf("%#v\n", p)      //Type + value
	fmt.Printf("%T\n", p)       //Type
	fmt.Printf("t: %t\n", true) //
	fmt.Printf("d: %d\n", 123)  //digital
	fmt.Printf("b: %b\n", 14)   //binary
	fmt.Printf("c: %c\n", 33)
	fmt.Printf("x: %x\n", 456)
	fmt.Printf("f: %f\n", 78.9)
	fmt.Printf("e: %e\n", 123400000.0)
	fmt.Printf("E: %E\n", 123400000.0)
	fmt.Printf("s: %s\n", "\"string\"")
	fmt.Printf("q: %q\n", "\"string\"")
	fmt.Printf("x: %x\n", "hex this")
	fmt.Printf("p: %p\n", &p)
	fmt.Printf("6d : |%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

type point struct {
	x, y int
}
