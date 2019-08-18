package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	calStringLength()
	TraversingChar()
	getStrSlice()
	changeStr()
	formatStr()
}

//计算字符串长度
func calStringLength() {
	str1 := "hello"
	str2 := "你好"
	fmt.Println(len(str1))
	fmt.Println(len(str2))
	fmt.Println(utf8.RuneCountInString(str2)) //计算utf-8字符串的长度
	str3 := str2 + str1
	fmt.Println(utf8.RuneCountInString(str3))
}

//遍历字符串
func TraversingChar() {
	str := "狙击 start"
	for i := 0; i < len(str); i++ {
		fmt.Printf("ascii: %c %d\n", str[i], str[i])
	}
	for _, s := range str {
		fmt.Printf("unicode: %c %d\n", s, s)
	}
}

//获取字符串片段
func getStrSlice() {
	str := "we are win,we really win"
	comm := strings.Index(str, ",")
	fmt.Println(str[comm:])
	fmt.Println(str[comm+1:])
}

//修改字符串
func changeStr() {
	str := "heros never die"
	strBytes := []byte(str)
	for i := 5; i < 11; i++ {
		strBytes[i] = ' '
	}
	fmt.Println(string(strBytes))
}

//格式化输出
func formatStr() {
	var a int = 8
	var b int = 92
	title := fmt.Sprintf("已经采集了%v个药草，还需要%v个", a, b)
	fmt.Println(title)
	proFile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   100,
	}
	fmt.Printf("使用'%%+v' %+v\n", proFile)
	fmt.Printf("使用'%%#v' %#v\n", proFile)
	fmt.Printf("使用'%%T' %T\n", proFile)
}
