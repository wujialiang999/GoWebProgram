package main

import (
	"fmt"
)

//类型定义
type NewInt int

//类型别名
type IntAlias = int

//类型别名与类型定义的区别
func main() {
	var a NewInt
	fmt.Printf("a type:%T\n", a)
	var a2 IntAlias
	fmt.Printf("a2 type:%T\n", a2)
}
