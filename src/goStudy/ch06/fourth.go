package main

import (
	"fmt"
)

type wheel struct {
	Size int
}

type engine struct {
	Power int
	Type  string
}

type car struct {
	wheel
	engine
}

func main() {
	c := car{
		wheel{Size: 10},
		engine{Power: 18, Type: "hh"},
	}
	fmt.Printf("%+v\n", c)
}
