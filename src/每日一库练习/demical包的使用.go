package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

func main() {

	div := decimal.NewFromInt(19).Div(decimal.NewFromInt(3))
	fmt.Println("div : ", div) // 除法

	f, _ := div.Float64()                                        // 先转成浮点数
	sprintf, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64) // 保留两位小数
	fmt.Println(sprintf)

	s := div.String()
	fmt.Println("s : ", s)

	round := div.Round(2) // 四舍五入的    6.125  --> 6.13, 6.124 --> 6.12
	fmt.Println("四舍五入2位： ", round.String())

	i := div.Ceil().String() // 向上去整数 6.33  ---> 7
	fmt.Println("i : ", i)

	i2 := div.Floor().String() // 向下取整数  6.33 ---- > 6
	fmt.Println("i2: ", i2)

	down := div.RoundDown(2) // todo 无视后面的小数进行向上或者向下取
	fmt.Println(down)
	down1 := div.RoundCeil(2)
	fmt.Println(down1)

	sub := decimal.NewFromInt(9000).Div(decimal.NewFromInt(100)).Sub(decimal.NewFromInt(900).Div(decimal.NewFromInt(100)))
	fmt.Println(sub.String())
	f2, _ := sub.Div(decimal.NewFromInt(9000).Div(decimal.NewFromInt(100))).Mul(decimal.NewFromFloat(0.5)).Float64()
	fmt.Println(" float data: ", f2)

}
