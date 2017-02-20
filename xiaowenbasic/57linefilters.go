// 57linefilters
package xiaowenbasic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LineFiltersMain() {
	fmt.Println("Line Filters Main")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
