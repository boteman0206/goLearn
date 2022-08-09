package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"time"
)

/*
https://learnku.com/articles/49065 参考文档

https://github.com/kevinyan815/gocookbook/issues/27

令牌桶的思想：
	想象有一个木桶，以固定的速度往木桶里加入令牌，木桶满了则不再加入令牌。服务收到请求时尝试从木桶中取出一个令牌，
如果能够得到令牌则继续执行后续的业务逻辑；如果没有得到令牌，直接返回访问频率超限的错误码或页面等，不继续执行后续的业务逻辑
	特点：由于木桶内只要有令牌，请求就可以被处理，所以令牌桶算法可以支持突发流量。

同时由于往木桶添加令牌的速度是固定的，且木桶的容量有上限，所以单位时间内处理的请求书也能够得到控制，起到限流的目的。
假设加入令牌的速度为 1token/10ms，桶的容量为500，在请求比较的少的时候（小于每10毫秒1个请求）时，木桶可以先"攒"一些令牌（最多500个）。
当有突发流量时，一下把木桶内的令牌取空，也就是有500个在并发执行的业务逻辑，之后要等每10ms补充一个新的令牌才能接收一个新的请求。

而当我们在使用一些第三方的限流时：我们可以将令牌桶的容量设置为1，这样最多就攒了一个令牌，然后每秒生成令牌的数量自己定义类似这样  rate.NewLimiter(5, 1),令牌桶的容量为1，每秒生成5个令牌即可


*/
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
