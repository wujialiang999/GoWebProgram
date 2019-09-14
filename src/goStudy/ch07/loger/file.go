package loger

import (
	"errors"
	"fmt"
	"os"
)

//文件写入器
type fileWrite struct {
	file *os.File
}

func (f *fileWrite) SetFile(filename string) (err error) {
	//如果文件已经打开，关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}
	f.file, err = os.Create(filename)
	return err
}
func (f *fileWrite) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("file not created!")
	}
	str := fmt.Sprintf("%v\n", data)
	_, err := f.file.Write([]byte(str))
	return err
}

func NewFileWrite() *fileWrite {
	return &fileWrite{}
}
