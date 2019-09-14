package loger

import (
	"fmt"
	"os"
)

//命令行写入器
type consoleWriter struct {
}

func (f *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

func NewConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
