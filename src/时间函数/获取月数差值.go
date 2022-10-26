package main

import (
	"fmt"
	"time"
)

var DATE_TIME_LAYOUT = "2006-01-02 15:04:05"

/**
获取相差的月份数量和天数
*/
func main() {
	location, err := time.ParseInLocation(DATE_TIME_LAYOUT, "2021-10-26 11:20:30", time.Local)
	if err != nil {
		return
	}
	month, day, wholeMonth := GetAge(time.Now(), location)
	fmt.Println("====: ", month, day, wholeMonth)
}

func GetAge(t1, t2 time.Time) (month, day, isWholeMonth int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 0大于 1等于 2小于
	isWholeMonth = 0
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		isWholeMonth = 2
		monthInterval--
	}
	if d1 == d2 {
		isWholeMonth = 1
	}
	day = d1 - d2
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}
