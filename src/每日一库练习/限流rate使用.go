package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limiter := rate.NewLimiter(1, 1)

	fmt.Println(limiter)

	for {
		time.Sleep(100 * time.Millisecond)
		//fmt.Println(limiter.Burst())
		ok1 := limiter.AllowN(time.Now(), 1)
		//time.Sleep(100*time.Millisecond)
		ok2 := limiter.AllowN(time.Now(), 1)

		//reserve := limiter.Reserve()
		//delay := reserve.Delay()
		//ok := reserve.OK()
		//fmt.Println("delay:", delay, "ok: ", ok )
		//
		//if !reserve.OK() {
		//	fmt.Println("exit（）")
		//	return
		//}
		err := limiter.Wait(context.Background())
		fmt.Println("err: ", err)

		ok3 := limiter.AllowN(time.Now(), 1)
		err = limiter.Wait(context.Background())
		fmt.Println("err: ", err)
		ok4 := limiter.AllowN(time.Now(), 1)
		err = limiter.Wait(context.Background())
		fmt.Println("err: ", err)
		ok5 := limiter.AllowN(time.Now(), 1)

		err = limiter.Wait(context.Background())
		fmt.Println("err: ", err)
		ok6 := limiter.AllowN(time.Now(), 1)
		ok7 := limiter.AllowN(time.Now(), 1)
		ok8 := limiter.AllowN(time.Now(), 1)
		ok9 := limiter.AllowN(time.Now(), 1)
		ok10 := limiter.AllowN(time.Now(), 1)
		fmt.Println("is ok :", ok1, ok2, ok3, ok4, ok5, ok6, ok7, ok8, ok9, ok10)

	}
}

//

//func main() {
//	limiter := rate.NewLimiter(rate.Every(time.Millisecond * 31), 2)
//	//time.Sleep(time.Second)
//	for i := 0; i < 10; i++ {
//		var ok bool
//		if limiter.Allow() {
//			ok = true
//		}
//		time.Sleep(time.Millisecond * 20)
//		fmt.Println(ok, limiter.Burst())
//	}
//}
