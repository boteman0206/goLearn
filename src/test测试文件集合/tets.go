package main

import (
	"fmt"
	"os"
)

// 获取输入的参数
func severityFrom(args []string) string {
	var s string
	fmt.Println("args: ", args)
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}

func main() {
	//from := severityFrom(os.Args)
	//fmt.Println(from)

	var a float64 = 20.35
	var c float64 = 1.123
	b := a + c - 1

	fmt.Println(b)

	var PackingCost, GoodsPayTotal, sumFreightPrivilege float64
	GoodsPayTotal = 20.35
	PackingCost = 0.1
	RefundAmount := GoodsPayTotal + PackingCost + 0 - sumFreightPrivilege
	fmt.Println(RefundAmount)

}
