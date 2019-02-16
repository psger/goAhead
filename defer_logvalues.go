package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q)= %d, %v", s, n, err)
	}() // 定义匿名函数并直接调用 最后一对括号表示对函数的调用
	return 7, io.EOF
}

func main()  {
	func1("go")
}
