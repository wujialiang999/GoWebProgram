package main

import (
	"errors"
	"fmt"
)

//定义除数为零的错误
func main(){
	fmt.Println(div(1,0))
}

var errDividionByZero=errors.New("除数为零")

func div (dividend ,divisor int)(int,error){
	if divisor==0{
		return 0,errDividionByZero
	}
	return dividend/divisor,nil
}