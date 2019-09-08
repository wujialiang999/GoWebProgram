package main

import "fmt"

//延迟执行语句
func main(){
	fmt.Println("延迟执行开始:")
	defer fmt.Println(1)
	defer fmt.Println(2)
	//栈队列,先进后出
	fmt.Println("延迟执行结束。")
}


