package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	brnary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	fmt.Println(brnary)

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	fmt.Println(env)

	execErr := syscall.Exec(brnary, args, env)

	if execErr != nil {
		panic(execErr)
	}
}
