package main

import "regexp"

//宕机
func main() {
	panic("宕机了")
}

//运行所必需的资源缺失时触发宕机
func mustCompile(str string) *Regexp {
	regexpResult, err := regexp.Complile(str)
	if err != nil {
		panic(`regexp Compile(` + quote(str) + `):` + err.Error())
	}
	return regexpResult
}
