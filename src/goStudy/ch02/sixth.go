package main

import (
	"fmt"
)

//变量逃逸分析
//go run -gcflags "-m -l" sixth.go
func main() {
	var a int
	void()
	fmt.Println(a, dummy(0))
}

func dummy(b int) int {
	var c int
	c = b
	return c
}

//空函数
func void() {

}
