package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	names := MyStringList{
		"1.zzh",
		"2.zzh",
		"3.zzh",
		"6.zzh",
		"5.zzh",
		"4.zzh",
	}
	sort.Sort(names)
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
	fmt.Println("")
	// 简化为以下代码
	name_row := sort.StringSlice{
		"1.zzh",
		"2.zzh",
		"3.zzh",
		"6.zzh",
		"5.zzh",
		"4.zzh",
	}
	sort.Sort(name_row)
	for _, v := range name_row {
		fmt.Printf("%s\n", v)
	}
	fmt.Println("")
	//更简单的方法
	names2 := []string{
		"1.zzh",
		"2.zzh",
		"3.zzh",
		"6.zzh",
		"5.zzh",
		"4.zzh",
	}
	sort.Strings(names2)
	for _, v := range names2 {
		fmt.Printf("%s\n", v)
	}
}
