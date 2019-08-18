package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//从setting.ini配置文件中查询需要的值
func main() {
	fileName := "D:\\goProject\\src\\goStudy\\ch02\\setting.ini"
	result := getValue2(fileName, "remote \"origin\"", "url")
	fmt.Println(result)
}
func getValue2(filePath, expectSection, expectKey string) (result string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//当前读取的段的名字
	var sectionName string
	for {
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == "" {
			continue
		}
		if linestr[0] == ';' {
			continue
		}
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			sectionName = linestr[1 : len(linestr)-1]
		} else {
			if sectionName == expectSection {
				pair := strings.Split(linestr, "=")
				if len(pair) == 2 {
					key := strings.TrimSpace(pair[0])
					if key == expectKey {
						result = strings.TrimSpace(pair[1])
					}
				}
			}
		}
	}
	return result
}
