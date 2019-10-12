package main

import (
	"fmt"
	"goPachong/zzx/trafficcrawler"
)

func main() {
	cityCodeURL := "https://jiaotong.baidu.com/trafficindex/city/list"
	body, err := trafficcrawler.HTTPGet(cityCodeURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = trafficcrawler.CreateCitycodeCSV(body, "citycode.csv")
	if err != nil {
		return
	}
	citycodes, err := trafficcrawler.GetCitycode("citycode.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	// citycodes := []string{"210", "79"}
	citychan := make(chan string)
	for i := 0; i < len(citycodes); i++ {
		go trafficcrawler.CreateCitydataCSV(citycodes[i], citychan)
	}
	for i := 0; i < len(citycodes); i++ {
		fmt.Printf("代码为%s的城市保存完毕", <-citychan)
	}
}
