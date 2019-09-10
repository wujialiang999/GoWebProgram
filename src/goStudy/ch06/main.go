package main

import (
	"fmt"
)

func main() {
	p := NewPlayer(0.5)
	p.MoveTo(Vec{3, 1})
	for !p.IsArrived() {
		p.Update()
		fmt.Println(p.Pos())
	}
}
