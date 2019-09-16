package main

import (
	"goStudy/ch08/base"
	_ "goStudy/ch08/cls1"
	_ "goStudy/ch08/cls2"
)

func main() {
	c1 := base.Create("Class1")
	c1.Do()
	c2 := base.Create("Class2")
	c2.Do()
}
