package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	//os.Exit(0),不会答应exit status
	//os.Exit(3)
	//log.Fatalf("不会执行defer")
}
