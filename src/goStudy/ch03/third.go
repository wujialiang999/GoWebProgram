package main

import (
	"fmt"
	"sort"
	"sync"
)

//映射
func main() {
	createMap()
	returnResult()
	deleteMap()
	MapInSync()
}

//创建Map
func createMap() {
	fmt.Println("创建Map")
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])
	v := scene["route2"]
	fmt.Println(v)
	//另一种声明方式
	var newMap = map[string]string{
		"A": "angle",
		"B": "body",
		"C": "cat",
	}
	for k, v := range newMap {
		fmt.Printf("key=%v,value=%v\n", k, v)
	}
}

//返回结果进行排序
func returnResult() {
	fmt.Println("返回结果进行排序")
	var newMap = map[string]string{
		"B": "body",
		"A": "angle",
		"C": "cat",
	}
	var tempList []string
	for k, v := range newMap {
		fmt.Printf("key=%v,value=%v\n", k, v)
		tempList = append(tempList, v)
	}
	sort.Strings(tempList)
	fmt.Printf("result=%v", tempList)
}

//删除键值对并清空元素
func deleteMap() {
	fmt.Println("删除键值对并清空元素")
	var newMap = map[string]string{
		"B": "body",
		"A": "angle",
		"C": "cat",
	}
	delete(newMap, "B")
	for k, v := range newMap {
		fmt.Printf("key=%v,value=%v\n", k, v)
	}
	//清空map
	newMap = make(map[string]string)
}

//并发环境下使用Map
func MapInSync() {
	fmt.Println("并发环境下使用Map")
	var scene sync.Map
	scene.Store("A", "angle")
	scene.Store("B", "body")
	scene.Store("C", "cat")
	fmt.Println(scene.Load("A"))
	scene.Delete("C")
	scene.Range(func(k, v interface{}) bool {
		fmt.Printf("k=%v,v=%v\n", k, v)
		return true
	})
}
