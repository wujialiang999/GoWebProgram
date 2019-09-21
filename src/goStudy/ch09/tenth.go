package main

import (
	"fmt"
	"time"
)

func main() {
	//打点器，500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)
	//计时器两秒后触发
	stopper := time.NewTimer(time.Second * 2)
	var i int
	for {
		select {
		case <-stopper.C:
			fmt.Println("Stop")
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("Tick", i)
		}
	}
StopHere:
	fmt.Println("Done")
}
