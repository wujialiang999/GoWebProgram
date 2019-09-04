package main

import "fmt"

//匿名函数
func main() {
	//定义时调用匿名函数
	func(data int) {
		fmt.Println(data)
	}(100)
	//匿名函数赋值给变量
	f := func(data int) {
		fmt.Println(data)
	}
	f(500)
	//匿名函数做回调函数
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
