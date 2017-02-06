// 50randomnumbers

/*
	output:
			Random Numbers main
			81,87
			0.6645600532184904
			7.1885709359349015,7.123187485356329
			60,85
			5,87
			5,87
*/
package xiaowenbasic

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumbersMain() {
	fmt.Println("Random Numbers main")

	p := fmt.Print
	pl := fmt.Println

	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	pl()

	pl(rand.Float64())

	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	pl()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	pl()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	pl()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	p(r3.Intn(100), ",")
	p(r3.Intn(100))
	pl()
}
