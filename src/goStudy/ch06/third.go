package main

import (
	"fmt"
)

type flying struct {
}

type walking struct {
}

func (f *flying) fly() {
	fmt.Println("flying")
}

func (w *walking) walk() {
	fmt.Println("walking")
}

type human struct {
	walking
}

type bird struct {
	walking
	flying
}

func main() {
	b := new(bird)
	fmt.Println("bird:")
	b.walk()
	b.fly()
	h := new(human)
	fmt.Println("human:")
	h.walk()
}
