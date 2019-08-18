package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

//生成sin函数的图像的灰度图
func main() {
	const size = 300
	pic := image.NewGray(image.Rect(0, 0, size, size)) //图像大小300*300像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255}) //图像背景设置为白色
		}
	}
	//从零到最大像素生成x坐标
	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{0})
	}
	file, err := os.Create("D:\\goProject\\src\\goStudy\\ch02\\sin.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, pic)

	file.Close()
}
