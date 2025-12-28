package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	fmt.Println([]byte(data)) //[97 98 99 49 50 51 33 63 36 42 38 40 41 39 45 61 64 126]

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(sEnc) //YWJjMTIzIT8kKiYoKSctPUB+

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)

	fmt.Println(sDec) //[97 98 99 49 50 51 33 63 36 42 38 40 41 39 45 61 64 126]

	fmt.Println(string(sDec)) //abc123!?$*&()'-=@~

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))

	fmt.Println(uEnc) //YWJjMTIzIT8kKiYoKSctPUB-

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)

	fmt.Println(sDec) //[97 98 99 49 50 51 33 63 36 42 38 40 41 39 45 61 64 126]

	fmt.Println(string(uDec)) //abc123!?$*&()'-=@~
}
