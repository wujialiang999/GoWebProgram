package main

import "fmt"

func main() {
	// 保存值
	var any interface{}
	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = false
	fmt.Println(any)
	// 取值
	var inta int = 1
	var INTa interface{} = inta
	// var INTb int =INTa 错误
	var INTb int = INTa.(int)
	fmt.Println(INTb)
	//比较
	var bja interface{} = 100
	var bjb interface{} = "hello"
	fmt.Println(bja == bjb)
	//map、切片不可比较
}
