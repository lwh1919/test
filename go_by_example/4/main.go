package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	dataOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataOut))

	grepHelloCmd := exec.Command("grep", "Hello")

	grepIn, _ := grepHelloCmd.StdinPipe()
	grepOut, _ := grepHelloCmd.StdoutPipe()
	grepHelloCmd.Start()
	grepIn.Write([]byte("Hello World\n"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepHelloCmd.Wait()

	fmt.Println("> grep Hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	dataOut, err = lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls ")
	fmt.Println(string(dataOut))
}
