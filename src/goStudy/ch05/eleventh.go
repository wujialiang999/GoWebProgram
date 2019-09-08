package main

import(
	"bytes"
	"fmt"
)
//获得可变参数类型
func main(){
	fmt.Println(printTypeValue("str",12,true))
}

func printTypeValue(slist ...interface{})string{
	var b bytes.Buffer
	for _,s:=range slist{
		str:=fmt.Sprintf("%v",s)
		var typeString string
		switch s.(type){
		case bool:
			typeString="bool"
		case string:
			typeString="string"
		case int:
			typeString="int"
		}
		b.WriteString("value: ")
		b.WriteString(str)
		b.WriteString("type: ")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}


