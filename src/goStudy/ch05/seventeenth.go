package main

import "fmt"

//在解析中使用自定义错误
func main(){
	var e error;
	//创建一个错误示例
	e=newParaseError("main.go",100)
	fmt.Println(e.Error())
	switch detail:=e.(type) {
	case *parseError:
		fmt.Printf("文件名:%s,行数:%d\n",detail.Filename,detail.Line)
	default:
		fmt.Println("其他类型错误")
	}
}

type parseError struct {
	Filename string
	Line int
}

func (e *parseError) Error() string{
	return fmt.Sprintf("%s:%d",e.Filename,e.Line)

}

func newParaseError(filename string,line int) error{
	return &parseError{filename,line}
}
