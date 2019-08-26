package main

import(
	"fmt"
)
//if语句
func main(){
	a:=10
	if a>10{
		fmt.Println("a>10")
	}else{
		fmt.Println("a<=10")
	}
	if result:=test();result!="aaa"{
		fmt.Println("error")
		return
	}
}
func test()(result string){
	return "dfd"
}
