package main

import (
	"fmt"
)

//取地址发生的变量逃逸
func main() {
	fmt.Println(dummy2())
}

type Data struct {
}

func dummy2() *Data {
	var c Data
	return &c
}
