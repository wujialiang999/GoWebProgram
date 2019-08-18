package main

import (
	"encoding/base64"
	"fmt"
)

//base64编码
func main() {
	mess := "I'm heros,do you believe me?渣渣辉"
	enMess := base64.StdEncoding.EncodeToString([]byte(mess))
	fmt.Println(enMess) //打印编码消息
	deMess, err := base64.StdEncoding.DecodeString(enMess)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(deMess))
	}
}
