package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"time"
)

func main() {

	l := rate.NewLimiter(5, 1)

	//fmt.Println("l.Burst() : ")
	time.Sleep(5 * time.Second)

	log.Println(l.Limit(), l.Burst())
	for i := 0; i < 100; i++ {
		//阻塞等待直到，取到一个token
		log.Println("before Wait")

		//返回需要等待多久才有新的token,这样就可以等待指定时间执行任务
		r := l.Reserve()
		fmt.Println("需要等待的时间： ", i, " 个数： ", r.Delay())

		log.Println("after Wait")
		//time.Sleep(r.Delay())
		//判断当前是否可以取到token
		//a := l.Allow()
		//log.Println("Allow:", a)
	}
}
