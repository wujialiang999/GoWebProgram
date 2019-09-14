package main

import (
	"fmt"
	"goStudy/ch07/loger"
)

func createLogger() *loger.Logger {
	l := loger.NewLogger()
	cw := loger.NewConsoleWriter()
	l.RegisterWriter(cw)
	fw := loger.NewFileWrite()
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Printf("%v\n", err)
	}
	l.RegisterWriter(fw)
	return l
}

func main() {
	l := createLogger()
	l.Log("hello")
}
