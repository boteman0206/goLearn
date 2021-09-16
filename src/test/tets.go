package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	newInt := big.NewInt(123)
	fmt.Println("newInt: ", newInt)

	t, err := time.Parse("2006-01-02", "2020-10-14")
	fmt.Println(t, err)
	t, err = time.Parse("2006-01-02 15:04:05", "2020-10-14 00:00:00")
	fmt.Println(t, err)

	//time.ParseInLocation()

	slice1 := []int{5, 23, 10, 2, 61, 33}

	ints := slice1[1:4:6]
	fmt.Println(ints, len(ints), cap(ints))

	mapWarehouseGoodsStocks := make(map[string]int32, 0)
	v := struct {
		Goodsid string
		Stock   int
	}{"乐视薯片", 18}
	mapWarehouseGoodsStocks[v.Goodsid] = mapWarehouseGoodsStocks[v.Goodsid] + int32(v.Stock)

	fmt.Println("mapWarehouseGoodsStocks :", mapWarehouseGoodsStocks)

	a := make([]int, 1, 10)
	//stringsMap := make(map[int]string)

	fmt.Println(a)

	sprintf := fmt.Sprintf("2,%s", "th090RJ")
	fmt.Println(sprintf)

	patternkeys := fmt.Sprintf("warehouse:%d:*", 5)
	fmt.Println(patternkeys)

	s := "asascs"
	d := "warehouse:" + s
	fmt.Println(d)

	unix := time.Now().Add(-1 * time.Minute * 5).Unix() // 1631758204
	fmt.Println(unix)
	i := time.Now().Unix() // 1631758504
	fmt.Println(i)

	timeDuration := time.Duration(-1*15) * time.Minute

	fmt.Println(timeDuration)
}
