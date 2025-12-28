package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bg := context.Background()
	todo := context.TODO()
	fmt.Println("background", bg)
	fmt.Println("TODO", todo)

	ctxCancel, cancel := context.WithCancel(bg)
	go func() {
		select {
		case <-ctxCancel.Done():
			fmt.Println("ctxCancel cancelled:", ctxCancel.Err())
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()

	ctxTimeout, cancelTimeout := context.WithTimeout(bg, time.Second)
	defer cancelTimeout()
	start := time.Now()
	select {
	case <-ctxTimeout.Done():
		fmt.Println("ctxTimeout cancelled:", ctxTimeout.Err(), " ", time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}

	deadline := time.Now().Add(5 * time.Second)
	ctxDeadline, cancelDeadline := context.WithDeadline(bg, deadline)
	defer cancelDeadline()
	select {
	case <-ctxDeadline.Done():
		fmt.Println("ctxDeadline cancelled:", ctxDeadline.Err(), " ", time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}

	type ctxKey string
	ctxVal := context.WithValue(bg, ctxKey("key"), "value")
	fmt.Println("ctxVal", ctxVal)

	parent, cancelParent := context.WithCancel(bg)
	child, cancelChild := context.WithCancel(parent)
	go func() {
		select {
		case <-child.Done():
			fmt.Println("ctxCancel cancelled:", child.Err(), " ", time.Since(start))
		}
	}()
	cancelParent()
	time.Sleep(4 * time.Second)

	cancelChild()
}
