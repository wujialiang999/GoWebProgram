package main

import (
	"fmt"
	"math"
)

func main() {
	//变量声明
	var (
		a int
		b int
		c int
		d int
	)
	a = 10
	b = 20
	c, d = jiaohuan(a, b)
	fmt.Println("a=", a, "b=", b, "a=", c, "b=", d)
	//匿名变量
	e, _ := jiaohuan(a, b)
	fmt.Println("e=", e)
	printPI()
}

//a,b变量值交换
func jiaohuan(a, b int) (c int, d int) {
	a, b = b, a
	c = a
	d = b
	return
}

//输出Pi的值
func printPI() {
	fmt.Printf("PI=%f\n", math.Pi)
	fmt.Printf("PI=%.2f\n", math.Pi)
	fmt.Printf("maxInt16=%d\n", math.MaxInt16)
	fmt.Printf("minInt16=%d\n", math.MinInt16)
}
