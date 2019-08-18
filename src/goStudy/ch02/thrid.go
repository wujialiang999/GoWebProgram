package main

import (
	"fmt"
)

func main() {
	//定义多行字符串
	duohang := `
func main(){
	fmt.Println("hello")
}`
	fmt.Printf(duohang + "\n")
	//uint8类型的字符串
	var a byte = 'a'
	//rune类型utf8字符串
	var b rune = '你'
	fmt.Printf("%c 的ASCII编码为：%d \t类型为：%T\n", a, a, a)
	fmt.Printf("%c 的UTF8编码为%d \t类型为：%T\n", b, b, b)
	//切片输出字符串
	c := make([]uint8, 3)
	c[0] = 1
	c[1] = 2
	c[2] = 3
	fmt.Printf("c[1]=%d\n", c[1])
	d := "hello world"
	fmt.Println(d[6:])
}
