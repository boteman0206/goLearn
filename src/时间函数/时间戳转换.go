package main

import (
	"fmt"
	"strconv"
	"time"
)

//13位时间戳转换
func TimeToUnix(e time.Time) int64 {
	timeUnix, _ := time.Parse("2006-01-02 15:04:05", e.Format("2006-01-02 15:04:05"))
	return timeUnix.UnixNano() / 1e6
}

//13位时间戳转换
func UnixToTime(e string) (datatime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}

/**
时间戳为10位的和13位的，需要注意
*/
func main() {

	//10位的直接使用自带函数即可 1664329390
	format := time.Unix(1664329390, 0).Format("2006-01-02 15:04:05")
	fmt.Println(format)

	toTime, err := UnixToTime("1664276821175")
	if err != nil {
		return
	}
	fmt.Println(toTime.Format("2006-01-02 15:04:05"))

}
