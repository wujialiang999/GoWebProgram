package main

import "os"

//延迟释放文件句柄
func main() {

}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil { //判断文件打开是否成功
		return 0
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		//延迟机制触发
		return 0
	}
	size := info.Size()
	//延迟机制触发
	return size
}
