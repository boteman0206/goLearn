package main

//  https://github.com/robfig/cron
//go get github.com/robfig/cron/v3@v3.0.0

import (
	"github.com/robfig/cron/v3"
)
import "fmt"

func testCron() {
	// todo 秒级别的
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/3 * * * * *", func() { // 不是标准的cron表达式
		fmt.Println("3 秒执行一次。。。。")
	})
	c.Start()
}
func main() {

	// 正常的表达式
	c := cron.New() // 版本问题 五个参数的话是分开始， 6个参数的话是从秒开始3.0版本的就需要指定cron.WithSeconds()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })

	// 没3秒执行一次
	c.AddFunc("@every 3s", func() {
		fmt.Println("33333333333333333333333")
	})

	// 每分钟执行一次
	c.AddFunc("* * * * *", func() {
		fmt.Println("--------------------")
	})

	//go testCron()
	c.AddFunc("0 46 17 * * ?", func() { fmt.Println("执行") })
	c.Start()
	select {}
}
