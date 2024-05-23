package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 假设有一个整数数组 numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// 定义一个变量 start 用于表示每次取出元素的起始下标
	start := 0

	// 循环取出 10 个元素，直到取完整个数组
	for start < len(numbers) {
		// 计算本次取出的元素个数
		end := start + 2
		if end > len(numbers) {
			end = len(numbers)
		}

		// 从数组中取出 10 个元素，并进行处理
		batch := numbers[start:end]
		fmt.Println(batch) // 做一些操作，例如打印出来

		// 更新下一次取出元素的起始下标
		start = end
	}

	// t := time.Now().Unix() - 1699845546

	t1 := time.Unix(time.Now().Unix(), 0) // 时间戳1，2021-07-20 00:00:00
	t2 := time.Unix(1699845546, 0)        // 时间戳2，2021-07-30 00:00:00

	duration := t1.Sub(t2)
	hours := int(duration.Hours())
	fmt.Println("-----", hours)

	ty := DaysBetween(1689217382, time.Now().Unix())
	fmt.Println("0000", ty)

	t001 := GetHourDateStr(0)
	t002 := GetHourDateStr(-1)
	fmt.Println("t001:", t001, " t002:", t002)

	var jsonData string = "121312342"

	var data int64

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("解析JSON数据失败:", err)
		return
	}
	fmt.Println("----", data)
	getLocalTime()

	batchSize := 10
	numBatches := (101 + batchSize - 1) / batchSize
	fmt.Println("numBatches---- ", numBatches)
	// 每次插入 batchSize 条数据
	for i := 0; i < numBatches; i++ {
		// 计算当前批次的起始位置和结束位置
		start := i * batchSize
		end := (i + 1) * batchSize
		if end > 101 {
			end = 101
		}
		fmt.Println("=====start, ", start, end)
	}

	t11 := []int64{0}

	fmt.Println("9: ", t11[0:2])

}

var (
	FORMAT_HOUR_STR    = "2006010215"
	STD_FORMAT_DAY_STR = "2006-01-02"
	// 标准秒级格式字符串
	STD_FORMAT_SECOND_STR = "2006-01-02 15:04:05"
)

func getLocalTime() {
	PkTimePointSlice := []string{"00:00:00", "02:00:00", "12:30:00", "20:00:00", "22:00:00"}
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Printf("加载时区失败：%v\n", err)
		return
	}
	var nextTime int64
	date := time.Now().In(loc).Format(STD_FORMAT_DAY_STR)
	for _, tp := range PkTimePointSlice {
		startDate := date + fmt.Sprintf(" %s", tp)
		t, err := time.ParseInLocation(STD_FORMAT_SECOND_STR, startDate, loc)
		if err != nil {
			fmt.Printf("解析时间点失败：%v\n", err)
			return
		}
		timestamp := t.Unix()
		fmt.Printf("%s 对应的时间戳为：%d, now:%d\n", t.Format("2006-01-02 15:04:05"), timestamp, time.Now().Unix())
		fmt.Println("----:", t.Format("15:04"))
		if time.Now().Unix() < timestamp {
			nextTime = timestamp

		}
	}

	fmt.Println("nextTime: ", nextTime)
}

func GetHourDateStr(hour int) string {
	curHourUnix := GetSpecifyHourTime(hour, nil).Unix() // 获取当前小时的起始时间
	curHourStr := ConvertToTimeStr(curHourUnix, FORMAT_HOUR_STR, TokyoLocation)
	return curHourStr
}
func GetSpecifyHourTime(hour int, location *time.Location) time.Time {
	if location == nil {
		location = time.Local
	}
	curTime := time.Now().In(location).Add(time.Hour * time.Duration(hour))
	timeStr := curTime.Format(STD_FORMAT_DAY_STR)
	t, _ := time.ParseInLocation(STD_FORMAT_SECOND_STR, timeStr+fmt.Sprintf(" %d:00:00", curTime.Hour()), location)
	return t
}

func ConvertToTimeStr(timeStamp int64, format string, location *time.Location) string {
	var t = time.Unix(timeStamp, 0)
	if location == nil {
		location = time.Local
	}
	return t.In(location).Format(format)
}

func DaysBetween(t1, t2 int64) int {

	if t1 < t2 {
		t1, t2 = t2, t1
	}
	// 将时间戳转换为 time.Time 对象
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)

	// 计算相差的天数
	duration := time1.Sub(time2)
	fmt.Println("-----", duration.Hours())
	days := duration.Hours() / 24
	fmt.Println("days", days)
	day1 := int(duration.Hours() / 24)
	fmt.Println("day1", day1)

	if days > float64(day1) {
		return day1 + 1
	}
	return day1
}

func VersionCompare(version1 string, version2 string) int {
	arrVersion1 := strings.Split(version1, ".")
	arrVersion2 := strings.Split(version2, ".")
	lenVer1 := len(arrVersion1)
	lenVer2 := len(arrVersion2)
	lenVer := lenVer1
	if lenVer1 < lenVer2 {
		lenVer = lenVer2
	}
	for i := 0; i < lenVer; i++ {
		var intV1, intV2 int
		if i < lenVer1 {
			intV1, _ = strconv.Atoi(arrVersion1[i])
		}
		if i < lenVer2 {
			intV2, _ = strconv.Atoi(arrVersion2[i])
		}
		if intV1 < intV2 {
			return -1
		}
		if intV1 > intV2 {
			return 1
		}
	}
	return 0
}
func chunkSlice(slice []int64, chunkSize int) [][]int64 {
	sliceSize := len(slice)
	numChunks := (sliceSize + chunkSize - 1) / chunkSize
	result := make([][]int64, numChunks)
	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > sliceSize {
			end = sliceSize
		}
		result[i] = slice[start:end]
	}
	return result
}

var (
	TokyoLocation, _ = time.LoadLocation("Asia/Tokyo")
)

func GetSpecifyDayTimeInterval(location *time.Location, day int) (start int64, end int64) {
	STD_FORMAT_DAY_STR := "2006-01-02"
	STD_FORMAT_SECOND_STR := "2006-01-02 15:04:05"
	if location == nil {
		// 默认使用本地时区
		location = time.Local
	}
	date := time.Now().AddDate(0, 0, day).In(location).Format(STD_FORMAT_DAY_STR)
	startDate := date + " 00:00:00"
	startTime, _ := time.ParseInLocation(STD_FORMAT_SECOND_STR, startDate, location)
	startUnix := startTime.Unix()
	return startUnix, startUnix + 24*3600 - 1
}
