package main

import "fmt"
import "strings"
func main() {
    list:=[]string{
        "go abc",
        "go bcd",
        "go dfdw",
        "go fdfdrfejk",
    }
    chain:=[]func(string)string{
        removePrefix,
        strings.TrimSpace,
        strings.ToUpper,
    }
    StringProccess(list,chain)
    for_,str:=range list{
        fmt.println(str)
    }
}

func StringProccess(list []string, chain []func(string) string){
    for index,str:=range list{
        result:=str
        for _,proc := range chain{
            result =proc(result)
        }
        list[index]=result
    }
}
func removePrefix(str string)string{
    return strings.TrimPrefix(str,"go")
}
