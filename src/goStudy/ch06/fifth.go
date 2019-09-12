package main

import (
	"fmt"
)

type wheel2 struct {
	Size int
}

type car2 struct {
	wheel2
	engine2 struct {
		power   int
		carType string
	}
}

func main() {
	c := car2{
		wheel2{Size: 10},
		struct {
			power   int
			carType string
		}{power: 18, carType: "hh"},
	}
	fmt.Printf("%+v\n", c)
}
