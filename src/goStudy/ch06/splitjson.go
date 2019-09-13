package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32
	Resx, ResY int
}

type Battery struct {
	Capacity int
}

func genJSONData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen{Size: 5.5, Resx: 1920, ResY: 1080},
		Battery{Capacity: 2910},
		true,
	}
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	jsonData := genJSONData()
	fmt.Println(string(jsonData))
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	//反序列化
	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Printf("%+v\n", screenAndTouch)
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData, batteryAndTouch)
	fmt.Printf("%+v\n", batteryAndTouch)
}
