package main

import (
	"fmt"
	"github.com/disintegration/imaging"
)

func main() {

	img, err := imaging.Open("D:\\RpPet\\gitProject\\goLearn\\src\\每日一库练习\\图片压缩包\\430.jpg")
	if err != nil {
		fmt.Println("打开图片失败：", err.Error())
	}
	//img = imaging.CropAnchor(img, 500, 800, imaging.Center) // 直接从中间裁剪，没有缩放
	//fmt.Println(src)

	img = imaging.Resize(img, 800, 800, imaging.Lanczos) // 这个可以等比例缩放

	err = imaging.Save(img, "D:\\RpPet\\gitProject\\goLearn\\src\\每日一库练习\\图片压缩包\\out_example.jpg")
	fmt.Println(err)
}
