// 59commandlineflags
/*
	命令行标记
	cmd op:

		命令操作方式及结果如下：

		$ ./command-line-flags -word=opt
		word: opt
		numb: 42
		fork: false
		svar: bar
		tail: []

		$ ./command-line-flags -word=opt a1 a2 a3
		word: opt
		...
		tail: [a1 a2 a3]

		$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
		word: opt
		numb: 42
		fork: false
		svar: bar
		tail: [a1 a2 a3 -numb=7]

		$ ./command-line-flags -h
		Usage of ./command-line-flags:
		  -fork=false: a bool
		  -numb=42: an int
		  -svar="bar": a string var
		  -word="foo": a string

		$ ./command-line-flags -wat
		flag provided but not defined: -wat
		Usage of ./command-line-flags:
		...

*/
package xiaowenbasic

import (
	"flag"
	"fmt"
)

func CommandLineFlagsMain() {
	fmt.Println("Command Line Flags main")

	pln := fmt.Println

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	pln("word:", *wordPtr)
	pln("numb:", *numbPtr)
	pln("fork:", *boolPtr)
	pln("svar:", svar)
	pln("tail:", flag.Args())
}
