package main

import (
	"fmt"
	// "log"
	"math"
	"strconv"
	"strings"
	"time"
	// "unicode/utf8"
)

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

func NowTimeInterval(db_time int64) float64 {
	now := time.Now().Unix()
	// 计算时间差
	duration := time.Unix(now, 0).Sub(time.Unix(db_time, 0))
	// 将时间差转换为小时数
	return duration.Hours()

}

func GetWeekTimeInterval(location *time.Location, week int) (start int64, end int64) {
	if location == nil {
		location = time.Local
	}
	now := time.Now().In(location)
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	day := week * 7
	lastWeekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location).AddDate(0, 0, offset+day)
	lastWeekStartTm := lastWeekStart.Unix()
	return lastWeekStartTm, lastWeekStartTm + 7*24*60*60 - 1
}

var (
	// 东京时区（+9时区）
	TokyoLocation, _ = time.LoadLocation("Asia/Tokyo")
	// 天级格式字符串
	FORMAT_DAY_STR = "20060102"
)

func ConvertToTimeStr(timeStamp int64, format string, location *time.Location) string {
	var t = time.Unix(timeStamp, 0)
	if location == nil {
		location = time.Local
	}
	return t.In(location).Format(format)
}

func GetWeekDateStr(week int) (start string, end string) {
	weekStart, weekEnd := GetWeekTimeInterval(TokyoLocation, week)
	startDate := ConvertToTimeStr(weekStart, FORMAT_DAY_STR, TokyoLocation)
	endDate := ConvertToTimeStr(weekEnd, FORMAT_DAY_STR, TokyoLocation)
	return startDate, endDate
}

func main() {

	// TokyoLocation, _ := time.LoadLocation("Asia/Tokyo")

	// jpTime := time.Now().In(TokyoLocation)
	// hourNum := jpTime.Hour()

	// fmt.Println(hourNum)
	// t := VersionCompare("asdas", "7.2.3")
	// fmt.Println(t)
	hur := NowTimeInterval(1609459200)
	fmt.Println(hur)

	fmt.Printf("%f", float64(123))

	t1, t2 := GetWeekDateStr(0)
	fmt.Println("-----------: ", t1, t2)

	japanTime := time.Unix(1695049200, 0).In(TokyoLocation)
	dateString := japanTime.Format("20060102")
	fmt.Println("dateString == ", dateString)

	fmt.Println(10%10 + 1)

	fmt.Println("================")
	a := 0b01
	b := 0b10
	fmt.Println(a & b)

	fmt.Println("-1<0: ", -1 < 0)
	fmt.Println("0<=0: ", 0 < 0)

	fmt.Println("----", 153107586%100)

	t := VersionCompare("9.2.9", "9.3")
	fmt.Println("版本比较-----", t, t < 0)

	fmt.Println(math.MaxInt32, math.MaxUint32)

	var x1 map[string]string = nil
	fmt.Println("map----:", x1["name"])
	// x1["name"] = "jack"
	// var s11 []int64 = nil
	// fmt.Println(s11)

	s11 := "我..是神"
	runes := []rune(s11)

	fmt.Println("rune:", runes)
	runes[0] = rune(31070)
	fmt.Println("rune2:", string(runes))

	// log.Fatalln("Fatalln----")
	// log.Panicln("panic-----")

	bubble01 := bubuleTest([]int64{99, 23, 4, 73, 2, 34, 44})
	fmt.Println("bubble01: ", bubble01)
	quick001 := quickTest([]int64{99, 23, 4, 73, 2, 34, 44})
	fmt.Println("quick001: ", quick001)

}

func bubuleTest(list []int64) []int64 {

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j+1] < list[j] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}

	return list
}

func quickTest(list []int64) []int64 {

	if len(list) <= 1 {
		return list
	}

	provid := list[0]
	left, right := []int64{}, []int64{}

	for _, v := range list[1:] {
		if v <= provid {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickTest(left)
	right = quickTest(right)

	return append(append(left, provid), right...)
}
