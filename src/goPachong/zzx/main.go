package main

import (
	"fmt"
	"goPachong\zzx\httputils"
)

func main() {
	//url := "https://jiaotong.baidu.com/trafficindex/city/list"
	//zzx.HTTPGet(url)
	//zzx.CreateCSV()
	resultCode := zzx.GetCSV()
	for i := 0; i < len(resultCode); i++ {
		fmt.Println("http://jiaotong.baidu.com/top/report/?citycode=" + resultCode[i])
	}

}

