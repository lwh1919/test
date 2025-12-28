package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 创建一个测试文件
	filePath := "/tmp/dat"
	content := []byte("hello\ngo\n")
	err := os.WriteFile(filePath, content, 0644)
	check(err)

	// -------------------------------
	// 1️⃣ 直接读取整个文件
	// -------------------------------
	dat, err := os.ReadFile(filePath)
	check(err)
	fmt.Println("ReadFile content:")
	fmt.Print(string(dat)) // 输出整个文件内容

	// -------------------------------
	// 2️⃣ 打开文件并分块读取
	// -------------------------------
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	b1 := make([]byte, 5) // 读取前 5 个字节
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("\nFile.Read: %d bytes: %s\n", n1, string(b1[:n1]))

	// -------------------------------
	// 3️⃣ 使用 Seek 定位文件指针
	// -------------------------------
	offset, err := f.Seek(6, 0) // 移动到第 6 个字节
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("Seek + Read: %d bytes @ %d: %s\n", n2, offset, string(b2[:n2]))

	// -------------------------------
	// 4️⃣ 使用 io.ReadAtLeast
	// -------------------------------
	_, err = f.Seek(6, 0) // 再次移动到第 6 个字节
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("ReadAtLeast: %d bytes: %s\n", n3, string(b3))

	// -------------------------------
	// 5️⃣ bufio 缓冲读取
	// -------------------------------
	_, err = f.Seek(0, 0) // 回到开头
	check(err)
	reader := bufio.NewReader(f)
	peekBytes, err := reader.Peek(5) // 查看前 5 个字节，不移动指针
	check(err)
	fmt.Printf("bufio Peek: %s\n", string(peekBytes))

	// 使用 ReadString 读取一行
	line, err := reader.ReadString('\n')
	check(err)
	fmt.Printf("bufio ReadString: %s", line)
}
