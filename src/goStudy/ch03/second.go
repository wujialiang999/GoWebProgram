package main

import "fmt"

//切片
func main() {
	fmt.Println("arrayTest1")
	arrayTest1()
	fmt.Println("arrayTest2")
	arrayTest2()
}

//指定范围生成切片
func arrayTest1() {
	//声明一个空切片
	var emptySlice = []int{}
	var building [30]int
	for i := 0; i < 30; i++ {
		building[i] = i + 1
	}
	fmt.Println(building[10:15])
	fmt.Println(building[20:])
	fmt.Println(building[:2])
	emptySlice = building[2:4]
	fmt.Println(emptySlice)
	//表示原有切片
	a := [3]int{1, 2, 3}
	fmt.Println(a[:])
	//重置切片
	fmt.Println(a[0:0])
}

//切片的增加、复制、删除操作
func arrayTest2() {
	a := make([]int, 2, 10)
	for i := 0; i < 5; i++ {
		a = append(a, i)
		fmt.Printf("a=%v i=%v\n", a, i)
	}
	//添加多个元素
	a = append(a, 5, 6)
	//添加切片
	b := make([]int, 3, 3)
	a = append(a, b...)
	fmt.Printf("a=%v\n", a)
	//复制
	var total int = 1000
	srcData := make([]int, total)
	for i := 0; i < total; i++ {
		srcData[i] = i + 1
	}
	refData := srcData
	copyData := make([]int, total)
	copy(copyData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copyData[0], copyData[total-1])
	//复制切片的从4到6
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Println(copyData[i])
	}
	//切片中删除元素
	seq := []string{"a", "b", "c", "d"}
	fmt.Println(seq)
	//指定删除位置
	index := 2
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}
