package main

import (
	"fmt"
)

//可变参数传递
func main() {
	print("fdfdf", "sdgg", "顶顶顶", 123)
}

func print(slist ...interface{}) {
	rawPrint(slist...)
}

func rawPrint(slist ...interface{}) {
	for _, s := range slist {
		fmt.Println(s)
	}
}
