package main

import (
	"bytes"
	"fmt"
)

//可变参数
func main() {
fmt.Println(joinStrings("a","b","c"))
fmt.Println(joinStrings("fd","fdsf"))
}
func joinStrings(slist ...string) string {
	var b bytes.Buffer
	for _, s := range slist {
		b.WriteString(s)
	}
	return b.String()
}
