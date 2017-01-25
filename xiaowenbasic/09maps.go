// 09maps
package xiaowenbasic

import (
	"fmt"
)

func MapMain() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	x := make(map[int]string)
	x[1] = "one"
	x[2] = "two"
	fmt.Println("x.map:", x)

	for k, v := range m {
		fmt.Println(">>>>>>", k, v, "\n")
	}

	for i := 1; i <= len(x); i++ {
		fmt.Println("<<<<<<", x[i])
	}

}
