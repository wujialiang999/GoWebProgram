package main

import "fmt"

//闭包
func main(){
	//简单使用
	str :="大家好"
	f:=func(){
		str="我是闭包"
		fmt.Println(str)
	}
	f()
	fmt.Println(str)
	//闭包的记忆效应
	accumulate:=Accumulate(1)
	fmt.Println(accumulate())
	fmt.Println(accumulate())
	fmt.Printf("指针为:%p\n",accumulate)

	accumulate2:=Accumulate(10)
	fmt.Println(accumulate2())
	fmt.Printf("指针为:%p\n",accumulate2)
}

func Accumulate(value int) func() int{
	return func() int {
		value++
		return value
	}
}

