package main

import "fmt"
import "runtime"

func main() {

	fmt.Println("CPU 核心数为:", runtime.NumCPU())
}
