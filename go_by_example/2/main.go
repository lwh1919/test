package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//old
	//signs := make(chan os.Signal, 1)
	//
	//signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)
	//
	//done := make(chan struct{})
	//
	//go func() {
	//	sig := <-signs
	//	fmt.Println()
	//	fmt.Println(sig)
	//	done <- struct{}{}
	//}()
	//<-done
	//fmt.Println("exit")

	//new
	ctx, stop := signal.NotifyContext( //当任一信号到达时自动调用 cancelctx.Done() 被关闭
		context.Background(), //以整个进程生命周期作为起点
		os.Interrupt,         //跨平台的 SIGINT
		syscall.SIGTERM)      // Unix 标准终止信号
	defer stop() //取消监听 + 资源清理
	<-ctx.Done()
	fmt.Println("\nexit")
}
