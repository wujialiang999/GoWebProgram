package main

import (
	"fmt"
)

var eventByName = make(map[string][]func(interface{}))

//注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	list := eventByName[name]
	list = append(list, callback)
	eventByName[name] = list
}

//调用事件
func CallEvent(name string, param interface{}) {
	list := eventByName[name]
	//遍历这个事件的所有回调
	for _, callback := range list {
		//传入参数调用回调
		callback(param)
	}
}

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

//全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {
	a := new(Actor)
	// RegisterEvent("OnSkill",GlobalEvent)
	RegisterEvent("OnSkill", a.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)
	CallEvent("OnSkill", 100) //会按照注册顺序执行
}
