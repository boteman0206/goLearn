package main

import (
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//discountPrice := decimal.NewFromInt(int64(977+300)).
	//	Mul(decimal.NewFromInt(int64(50))).
	//	DivRound(decimal.NewFromInt(100), 0).IntPart()
	//
	//fmt.Println(discountPrice)
	//
	//split := strings.Split("", ",")
	//fmt.Println(len(split), split, split[0])
	//
	//lastTime := time.Now()
	//for i:= 0; i<100 ; i ++  {
	//	subSecond := math.Ceil(6 - time.Now().Sub(lastTime).Seconds())
	//	subSecond = math.Ceil(subSecond)
	//	if subSecond > 0 {
	//		time.Sleep(time.Duration(subSecond) * time.Second)
	//	}
	//	lastTime = time.Now()
	//
	//	fmt.Println("last time : ", lastTime)
	//}

	//TablePrefix := ""
	//repeat := strings.Repeat(TablePrefix+" ", 10)
	//fmt.Println("10 次数：", repeat, len(repeat))
	//
	//
	//TablePrefix_gj := "gj_"
	//repeat_gj := strings.Repeat(TablePrefix_gj+" ", 10)
	//fmt.Println("10 次数：", repeat_gj, len(repeat_gj))
	//
	//TablePrefix_channel := "channel_"
	//repeat_channel := strings.Repeat(TablePrefix_channel+" ", 10)
	//fmt.Println("10 次数：", repeat_channel, len(repeat_channel))

	//key := "productcenter:auto:barcode:"
	source := rand.NewSource(time.Now().UnixNano())
	barcodeStr := time.Now().Format("20060102150405")
	t, _ := time.Parse("20060102150405", barcodeStr)
	endExp := t.UnixNano() + 57600
	expInt := time.Duration(endExp - time.Now().UnixNano())
	fmt.Println("expInt = ", int64(expInt), "endExp", endExp)
	var barcodeNum1 int32
	for i := 0; i < 10000; i++ {
		for true {
			barcodeNum1 = rand.New(source).Int31n(1000000)
			if barcodeNum1 > 0 {
				break
			}
		}
		fmt.Println("66666666", barcodeNum1)
	}

	for {
		//if redisLock := redis.SetNX(key+"lock", barcodeStr, 10*time.Second).Val(); !redisLock {
		//	continue
		//}
		//barcodeNumStr := redis.Get(key + barcodeStr).Val()
		barcodeNumStr := "15123"
		barcodeNum := cast.ToInt32(barcodeNumStr)
		BarCode := fmt.Sprintf("%s%05d", barcodeStr, barcodeNum+1)
		//if in.Product.ProductType == 3 {
		//	in.Product.BarCode = fmt.Sprintf("%s%05d", barcodeStr, barcodeNum+1)
		//}
		//redis.Set(key+barcodeStr, barcodeNum+1, expInt*time.Microsecond)
		fmt.Println("barcode: ", BarCode, expInt*time.Microsecond, expInt*time.Second)
		break

	}
	//redis.Del(key + "lock")
	//source := rand.NewSource(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		sprintf := fmt.Sprintf("%08d", rand.New(source).Int31n(1000000))
		fmt.Println(sprintf)

	}

	//product := Product{}

	//product.lock.Lock()
	//
	//product.lock.Unlock()
	for i := 0; i <= 100; i++ {
		//test19(product, i)
	}
	fmt.Println("--------------")

	BarCode := fmt.Sprintf("%s%05d", barcodeStr, 00000)
	fmt.Println("00000000", BarCode)
	//barcodeNum := cast.ToInt32("00000")

	i := int64(expInt) + 1213321
	fmt.Println("00000", i)
	BarCode1 := fmt.Sprintf("%s%05d", barcodeStr, i)
	fmt.Println("00000000", BarCode1)

}

type Product struct {
	categoryNames []string
	lock          sync.Mutex
}

func test19(this Product, num int) {
	this.lock.Lock()

	defer this.lock.Unlock()
	time.Sleep(1 * time.Second)
	fmt.Println("===============", num)
}
