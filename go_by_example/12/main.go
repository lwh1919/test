package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func C(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample.txt")
	C(err)

	fmt.Println("Filename:", f.Name())
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	C(err)

	dname, err := os.MkdirTemp("", "sample")
	fmt.Println(dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "sample.txt")
	err = os.WriteFile(fname, []byte{}, 0666)
	C(err)
}
