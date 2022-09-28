package main

import (
	"fmt"
	"godemo/src/localRedis"
	"strconv"
	"time"
)

func main() {

	localRedis.SetupDB()

	conn := localRedis.GetRedisConn()
	fmt.Printf("%p", conn)
	//time.Now().Format(kit.TimeLayout)

	//for i:=0; i<100 ; i++  {
	//
	//	conn := localRedis.GetRedisConn()  // 0xc00005c060
	//	fmt.Printf("%p", conn)
	//	fmt.Println("\n")
	//	time.Sleep(10*time.Second)
	//}
	//1664329390
	format := time.Unix(1664276821175, 0).Format("2006-01-02 03:04:05")

	fmt.Println(format)

	unix := time.Now().Unix()
	fmt.Println("unix: ", unix)
	fmt.Println(time.Unix(1664329390, 0).Format("2006-01-02 03:04:05"))
	//time.Sleep(10 * time.Hour)

	toTime, err := UnixToTime("1664276821175")
	if err != nil {
		return
	}
	fmt.Println(toTime.Format("2006-01-02 03:04:05"))

}

func TimeToUnix(e time.Time) int64 {
	timeUnix, _ := time.Parse("2006-01-02 15:04:05", e.Format("2006-01-02 15:04:05"))
	return timeUnix.UnixNano() / 1e6
}

func UnixToTime(e string) (datatime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}
