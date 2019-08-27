package main

import (
	"fmt"
)

//switch分支
func main() {
	a := "hello"
	switch a {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default")
	}
	switch a {
	case "hello", "haha":
		fmt.Println("hello or haha")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default")
	}
	var i int = 11
	switch {
	case i > 10 && i < 20:
		fmt.Println("10<i<20")
		//此时switch后面不能在跟判断变量，判断的目标都没有了
	}
	c := "hello"
	switch {
	case c == "hello":
		fmt.Println("Hello")
		fallthrough //兼容c语言的代码
	case c != "":
		fmt.Println("World")
	}
}
