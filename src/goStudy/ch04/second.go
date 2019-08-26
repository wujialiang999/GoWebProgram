package main

import (
	"fmt"
)

//for循环
func main() {
	step := 2
	for ; step > 0; step-- {
		fmt.Printf("step=%d\n", step)
	}
	var i int
	for ; ; i++ {
		if i > 10 {
			fmt.Println("i=",i)
			break
		}
	}
	var j int
	for j<10{
		fmt.Println("j=",j)
		j++
	}
}
