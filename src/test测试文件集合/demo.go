package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

func main() {

	//split := strings.Split(":11011", ":")
	//fmt.Println("split : ", split)
	//for i, v := range split {
	//	fmt.Println("i : ", i, " v : ", v)
	//}

	var maxUpdateDate time.Time
	fmt.Println("this is : ", maxUpdateDate)

	var t1 time.Time

	fmt.Println("==时间判断=", maxUpdateDate == t1, t1.Equal(maxUpdateDate))

	now := time.Now()

	date := compareUpdateDate(maxUpdateDate, now)
	fmt.Println(date)

	reg, _ := regexp.Compile(`\d+`)
	skuId := reg.FindString("rpc error: code = Unknown desc = 仓库信息不存在，请核实仓库信息！编码：RP0158,")
	fmt.Println(skuId) //  0158
	errors.New("mypackage: invalid parameter")

	//json.Marshal()
}

func compareUpdateDate(t1, t2 time.Time) time.Time {
	if t1.Before(t2) {
		return t2
	}
	return t1
}
