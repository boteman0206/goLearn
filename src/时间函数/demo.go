package main

import (
	"fmt"
	"reflect"
	"time"
)

/**
time包
*/

func main() {
	now := time.Now()
	//时间和类型
	fmt.Println("now : ", now, reflect.TypeOf(now))

	unix := now.Unix()
	fmt.Println("unix : ", unix)

	year, month, day := now.Date()
	fmt.Println("year :", year, " month :", month, "  day :", day)

	fmt.Println(now.Month())
	fmt.Println(int(now.Month())) // 直接获取中文的月份
	fmt.Println(now.Year())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())

	//格式化日期和时间
	fmt.Println(now.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	// 构造数字是固定的
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	//go时间常量
	time.Sleep(100 * time.Millisecond) // 设置100毫秒
	//time.Sleep(1*time.Hour) // 设置1小时

	//随机数的种子
	fmt.Println("unix 秒:", now.Unix())
	fmt.Println("unix 纳秒:", now.UnixNano())

	i := time.Now().UnixNano()
	timeCounn()
	i2 := time.Now().UnixNano()
	fmt.Println(i, "  ", i2)
	fmt.Println("消费时间 : ", i2-i)

}

func timeCounn() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("timeCount...")
}
