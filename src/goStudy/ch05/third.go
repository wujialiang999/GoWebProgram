package main

import "fmt"

func fire(){
    fmt.println("fire");
}

func main() {
    var f func()
   f = fire
   f()
}
