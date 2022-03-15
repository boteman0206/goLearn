package main

import (
	"fmt"
	"time"
)

//https://github.com/0voice/Introduction-to-Golang/blob/main/Go%E8%87%AA%E5%B8%A6%E5%BA%93%E7%9A%84%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E.md

//4.1 日期格式 转 时间戳

func TimeStr2Time(fmtStr, valueStr, locStr string) time.Time {
	loc := time.Local
	if locStr != "" {
		loc, _ = time.LoadLocation(locStr) // 设置时区
	}
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	t, _ := time.ParseInLocation(fmtStr, valueStr, loc)
	return t
	//return t.Unix()  //todo 转成时间戳
}

//4.2 获取当前时间日期格式
func GetCurrentFormatStr(fmtStr string) string {
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	return time.Now().Format(fmtStr)
}

//时间戳 to 日期格式
func Sec2TimeStr(sec int64, fmtStr string) string {
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	return time.Unix(sec, 0).Format(fmtStr)
}

func main() {

	str2Time := TimeStr2Time("", "2021-04-02 12:11:12", "")
	fmt.Println("str转成日期格式：", str2Time.String())

	str := GetCurrentFormatStr("")
	fmt.Println("当前时间str： ", str)

	timeStr := Sec2TimeStr(time.Now().Unix(), "")
	fmt.Println("时间戳转日期格式：", timeStr)

}
