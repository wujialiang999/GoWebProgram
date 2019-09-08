package main

import (
	"fmt"
	"runtime"
)

//宕机恢复
func main() {
	fmt.Println("运行前:")
	ProtectRun(func() {
		fmt.Println("手动宕机前：")
		panic(&panicContext{
			"手动触发宕机",
		})
		fmt.Println("手动宕机后")
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}

//崩溃时传递的上下文信息
type panicContext struct {
	function string
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("运行时错误:", err)
		default:
			fmt.Println("错误:", err)
		}
	}()
	entry()
}
