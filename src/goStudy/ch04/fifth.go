package main

import (
	"fmt"
)

//goto语句
func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				fmt.Println("j=", j)
				goto BreakFor
			}
		}
	}
	return
BreakFor:
	fmt.Println("Done")
	//统一错误处理
	/*
		err:=firstError
		if err!= nil{
			goto Exit
		}
		err=secondError
		if err!= nil{
			goto Exit
		}
		Exit:
			exit()//退出程序
	*/
	//跳出指定循环
BreakFor2:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				fmt.Println("j=", j, "i=", i)
				break BreakFor2 //跳出两个for循环
			}
		}
	}
	//继续下一下循环
jixuFor:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 1 {
				fmt.Println("j=", j, "i=", i)
				continue jixuFor //跳出当前for循环
			}
		}
	}
}
