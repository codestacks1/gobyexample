// 54base64encoding
/*
	output:
			Base64 Encoding main
			YWJjMTIzIT8kKiYoKSctPUB+
			abc123!?$*&()'-=@~

			YWJjMTIzIT8kKiYoKSctPUB-
			abc123!?$*&()'-=@~
*/

package xiaowenbasic

import (
	b64 "encoding/base64"
	"fmt"
)

func Base64EncodingMain() {
	fmt.Println("Base64 Encoding main")

	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
