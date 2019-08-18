package main

import "fmt"

//枚举常量-字符串类型
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

//枚举常量-数值类型
type Weapon int

const (
	Arrow Weapon = iota
	Arrow1
	Arrow2
	Arrow3
	Arrow4
)

//枚举
func main() {
	enumInt()
	enumString()
}

func enumInt() {
	fmt.Println(Arrow, Arrow1, Arrow2, Arrow3, Arrow4)
	var enumTest Weapon = Arrow1
	fmt.Println(enumTest)
}
func enumString() {
	fmt.Printf("%s %d", CPU, GPU)
}

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
