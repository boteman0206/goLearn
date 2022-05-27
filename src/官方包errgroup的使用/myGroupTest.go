package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

type myGroup struct {
	group  errgroup.Group
	errMap sync.Map
}

func main() {
	var g myGroup

	// 启动第一个子任务,它执行成功
	g.group.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		// return errors.New("failed to exec #1")
		return nil
	})
	// 启动第二个子任务，它执行失败
	g.group.Go(func() error {
		time.Sleep(10 * time.Second)
		g.errMap.Store("fun2", "第2个方法返回的异常")
		//return errors.New("failed to exec #2")
		return nil
	})

	// 启动第三个子任务，它执行成功
	g.group.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #3")
		g.errMap.Store("fun3", "第三个方法返回的异常")

		return nil
	})
	// 等待三个任务都完成
	if err := g.group.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}

	g.errMap.Range(func(key, value any) bool {
		fmt.Println(key)
		fmt.Println(value)
		return true
	})

}
