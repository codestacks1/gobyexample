// 53sha1 hashes
/*
	output:
			SHA1 Hashes
			sha1 this string
			cf23df2207d99a74fbe169e3eba035e633b65d94
*/

package xiaowenbasic

import (
	"crypto/sha1"
	"fmt"
)

func SHA1HashesMain() {
	fmt.Println("SHA1 Hashes")

	s := "sha1 this string"
	h := sha1.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
