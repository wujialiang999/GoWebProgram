package main

import (
	"container/list"
	"fmt"
)

//列表
func main() {
	createList()
}

//列表声明、插入和遍历
func createList() {
	newList := list.New()
	var newList2 list.List
	newList.PushBack("后1")
	newList.PushFront("前1")
	newList2.PushBack("后2")
	element := newList2.PushFront("前2")
	newList2.InsertAfter("插入后2", element)
	newList2.InsertBefore("插入前2", element)
	newList2.Remove(element)
	fmt.Println("NewList:")
	for i := newList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("NewList2:")
	for i := newList2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
