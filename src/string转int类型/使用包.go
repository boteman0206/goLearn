package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strconv"
)

func main() {

	atoi, err := strconv.Atoi("178")
	if err != nil {
		return
	}

	fmt.Println(atoi)

	toInt32 := cast.ToInt32("90")
	fmt.Println(toInt32)

	parseInt, err := strconv.ParseInt("870", 0, 0)
	if err != nil {
		return
	}

	fmt.Println(parseInt)

	fmt.Println(strconv.ParseInt("-12", 10, 0))
	//转换成十进制int64strconv.ParseInt("2345",10,64)
	//转换成八进制int32strconv.ParseInt("0xFF",0,32)
	//转换成16进制int64strconv.ParseInt("FF",16,64)

}
