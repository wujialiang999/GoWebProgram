package trafficcrawler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//CityCodeStruct 接受城市编码的结构
type CityCodeStruct struct {
	Status int `json:"status"`
	Data   struct {
		List []struct {
			Time         string  `json:"time"`
			Citycode     string  `json:"citycode"`
			Cityname     string  `json:"cityname"`
			Index        string  `json:"index"`
			LastIndex    string  `json:"last_index"`
			IndexLevel   int     `json:"index_level"`
			Speed        string  `json:"speed"`
			CityCoords   string  `json:"city_coords"`
			Provincecode int     `json:"provincecode"`
			Provincename string  `json:"provincename"`
			WeekRate     float64 `json:"weekRate"`
		} `json:"list"`
		Datetime string `json:"datetime"`
	} `json:"data"`
	Message interface{} `json:"message"`
}

//CityDataStruct 每个城市的JSON反序列化
type CityDataStruct struct {
	Status int `json:"status"`
	Data   struct {
		List []struct {
			Index interface{} `json:"index"`
			Speed string      `json:"speed"`
			Time  string      `json:"time"`
		} `json:"list"`
	} `json:"data"`
	Message interface{} `json:"message"`
}

//HTTPGet 获取Get请求得到的数据
func HTTPGet(geturl string) (result []byte, err error) {
	resp, err := http.Get(geturl)
	if err != nil {
		fmt.Println("GET请求失败:", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取网页内容失败:", err)
		return nil, err
	}
	return body, nil
}

//CreateCitycodeCSV 创建保存城市编码的csv
func CreateCitycodeCSV(body []byte, csvname string) (err error) {
	mydata := CityCodeStruct{}
	if err = json.Unmarshal(body, &mydata); err != nil {
		fmt.Println(err)
		return err
	} else {
		file, err := os.Create(csvname)
		//file, err := os.OpenFile("text.csv", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer file.Close()
		file.WriteString("\xEF\xBB\xBF")
		w := csv.NewWriter(file)
		var data [][]string
		fmt.Println("success")
		for i := 0; i < len(mydata.Data.List); i++ {
			mydata := []string{mydata.Data.List[i].Time, mydata.Data.List[i].Citycode, mydata.Data.List[i].Cityname}
			data = append(data, mydata)
		}
		w.WriteAll(data) //写入数据
		w.Flush()
		fmt.Println("success")
	}
	return nil
}

//GetCitycode 通过csv获取城市编码
func GetCitycode(csvname string) (citycodes []string, err error) {
	fs, err := os.Open(csvname)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			fmt.Printf("can not read, err is %+v\n", err)
		}
		if err == io.EOF {
			break
		}
		citycodes = append(citycodes, row[1])
		//fmt.Println(row)
	}
	return citycodes, nil
}

//CreateCitydataCSV 创建每个城市的csv数据
func CreateCitydataCSV(code string, citychan chan string) (err error) {
	dayDataURL := "https://jiaotong.baidu.com/trafficindex/city/curve?cityCode=" + code + "&type=minute"
	body, err := HTTPGet(dayDataURL)
	if err != nil {
		citychan <- code
		return err
	}
	mydata := CityDataStruct{}
	if err = json.Unmarshal(body, &mydata); err != nil {
		citychan <- code
		return err
	}
	for i := 0; i < len(mydata.Data.List); i++ {
		fmt.Printf("%v\n", mydata.Data.List[i].Index)
		fmt.Printf("%v\n", mydata.Data.List[i].Speed)
		fmt.Printf("%v\n", mydata.Data.List[i].Time)
	}
	citychan <- code
	return nil
}
