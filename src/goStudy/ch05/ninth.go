package main

import "fmt"

//闭包生成器
func main() {
	gen := playGen("勇士")
	name, hp := gen()
	fmt.Printf("name=%v,hp=%v", name, hp)
}

func playGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}
