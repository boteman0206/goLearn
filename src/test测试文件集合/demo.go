package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var DATE_TIME_LAYOUT = "2006-01-02 15:04:05"
var DATE_LAYOUT = "2006-01-02"

func main() {

	var ch = make(chan int, 10)
	go func() {
		for true {
			time.Sleep(10 * time.Second)
			ch <- 1
		}
	}()
	for i := range (<-chan int)(ch) {
		fmt.Println(i)
	}
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

func FilterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		}
	}
	return new_content
}

func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}
