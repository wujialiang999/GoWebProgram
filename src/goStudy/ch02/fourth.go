package main

import "fmt"

//指针学习
func main() {
	qudizhi()
	getValue()
	x, y := 1, 2
	fmt.Println("x=", x, "y=", y)
	swap(&x, &y)
	fmt.Println("x=", x, "y=", y)
}

//输出变量的地址
func qudizhi() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p  %p\n", &cat, &str)
}

//从指针获取指向指针的值
func getValue() {
	var str = "address"
	ptr := &str
	fmt.Println("str = ", str)
	fmt.Println("ptr's value =", *ptr)
	*ptr = "address change"
	fmt.Println("*ptr = address change")
	fmt.Println("str's value =", str)
}

//试用指针修改值
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
