package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	timer2 := time.NewTimer(time.Second * 1)

	<-timer1.C
	fmt.Println("timer1:", timer1)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	//stop2 := timer2.Stop()
	//if stop2 {
	//	fmt.Println("timer2 stopped")
	//}
	time.Sleep(time.Second * 1)
}
