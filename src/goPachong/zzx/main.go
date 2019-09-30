package main

import (
	"fmt"
	"goPachong/zzx/httputils"
)

func main() {
	fmt.Println("开始工作")
	url := "https://jiaotong.baidu.com/trafficindex/city/list"
	httputils.HTTPGet(url)
}
