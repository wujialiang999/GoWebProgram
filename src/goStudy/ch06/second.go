package main

import (
	"fmt"
)

type BasicColor struct {
	R, G, B float32
}
type Color struct {
	Basic BasicColor
	Alpha float32
}
type Color2 struct {
	BasicColor
	Alpha float32
}

func main() {
	var c Color
	c.Basic.R = 1
	c.Basic.G = 1
	c.Basic.B = 0
	c.Alpha = 1
	fmt.Printf("%+v\n", c)
	var c2 Color2
	c2.R = 1
	c2.G = 1
	c2.B = 1
	c2.Alpha = 1
	fmt.Printf("%+v\n", c2)
}
