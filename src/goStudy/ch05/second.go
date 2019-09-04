package main

import "fmt"

func main() {
    in:=Data{
        complax:[]int{1,2,3},
        instance:InnerData{
            5
        },
        ptr:&InnerData
    }
    fmt.Printf("in value:%+v\n",in)
    fmt.Printf("in value:%p\n",&in)
    out:=passByValue(in)
    fmt.Printf("out value:%+v\n",out)
    fmt.Printf("out value:%p\n",&out)
}
type InnerData{
    a int
}
type Data struct{
    complax []int,
    instance InnerData,
    ptr *InnerData
}
func passByValue(inFunc Data) Data{
    fmt.Printf("in value:%+v\n",inFunc)
    fmt.Printf("in value:%p\n",&inFunc)
    return inFunc
}
