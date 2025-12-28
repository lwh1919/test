package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	//recover 只能捕获“同一个 goroutine 中、在 defer 调用链上的 panic
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
