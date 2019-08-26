package main

import (
	"fmt"
)

//for range循环
func main() {
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key=%d value=%v\n", key, value)
	}
	var str = "大家好,haha"
	for key, value := range str {
		fmt.Printf("key=%d value=%v code=0x%x\n", key, value, value)
	}
	m := map[int](string){
		1: "one",
		2: "two",
		3: "three",
	}
	for key, value := range m {
		fmt.Printf("key=%d value=%v\n", key, value)
	}
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		c <- 5
		close(c)
	}()
	for value := range c {
		fmt.Printf("channel value=%v\n", value)
	}
}
