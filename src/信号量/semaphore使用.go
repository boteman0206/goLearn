package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"runtime"
	"time"
)

func queryDataWithSizeN(batchSize int) ([]int, error) {
	fmt.Println("queryDataWithSizeN run ...")
	return []int{1, 2, 3, 4, 5, 6}, nil
}
func useSemaphore() {
	var concurrentNum int64 = 10
	var weight int64 = 1
	var batchSize int = 50
	s := semaphore.NewWeighted(concurrentNum)
	for {
		data, _ := queryDataWithSizeN(batchSize)
		if len(data) == 0 {
			fmt.Println("End of all data")
			break
		}

		for _, item := range data {
			s.Acquire(context.Background(), weight)
			go func(i int) {
				fmt.Println("do something: ", i)
				time.Sleep(1 * time.Second)
				//doSomething(i)
				s.Release(weight)
			}(item)
		}

	}
}

func main() {
	//useSemaphore()
	//gomaxprocs := runtime.GOMAXPROCS(0) // 指定以1核运算
	gomaxprocs1 := runtime.GOMAXPROCS(1)
	//gomaxprocs2 := runtime.GOMAXPROCS(2)
	//fmt.Println(gomaxprocs)
	fmt.Println(gomaxprocs1)
	//fmt.Println(gomaxproc
	//
	//s2)

	//cpu := runtime.NumCPU()
	//fmt.Println(cpu)
	//po := runtime.GOMAXPROCS(cpu)
	//fmt.Println(po)

	for true {
		go func() {
			fmt.Println("hello ")
		}()
	}

}
