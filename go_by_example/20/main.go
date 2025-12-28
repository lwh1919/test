package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 cjaocjao"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println([]byte(s))
	fmt.Println(bs)

}
