package main

import (
	"fmt"
	"time"
)

func main() {
	a, _ := time.ParseInLocation("2006-01-02", "2017-09-01", time.Local)
	b, _ := time.ParseInLocation("2006-01-02", "2018-03-11", time.Local)
	d := a.Sub(b)

	fmt.Println(d.Hours() / 24)
}
