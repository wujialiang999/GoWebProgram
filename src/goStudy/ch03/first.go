package main

import (
	"fmt"
)

//创建数组，遍历数组
func main() {
	var team [3]string
	team[0] = "wjl"
	team[1] = "lc"
	team[2] = "中心"
	for k, v := range team {
		fmt.Println(k, v)
	}
}
