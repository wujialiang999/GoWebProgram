package second

/**
 *
 * 功能：百度贴吧并发爬虫
 * 作者:
 * 创建时间:2019年8月6日18:39:58
 */
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func SpiderPage(i int, page chan int) {
	url := "http://tieba.baidu.com/f?kw=gis&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := httpGet(url)
	if err != nil {
		fmt.Printf("HttpGet err ", err)
		return
	}
	//fmt.Println("result=",result)
	f, err := os.Create("第" + strconv.Itoa(i) + "页.html")
	if err != nil {
		fmt.Println("os err ", err)
		return
	}
	f.WriteString(result)
	f.Close()
	page <- i //与主go程协调工作
}

func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		if n == 0 {
			break
		}
		result += string(buf[:n])

	}
	return
}

func working(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页...\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)

	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个网页爬去完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Print("请输入爬取起始页\n")
	fmt.Scan(&start)
	fmt.Print("请输入爬取终止页\n")
	fmt.Scan(&end)
	working(start, end)
}
