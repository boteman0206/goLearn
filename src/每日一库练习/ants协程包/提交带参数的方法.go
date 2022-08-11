package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"sync/atomic"
)

var (
	sum  int32
	antP *ants.Pool
)

func init() {
	antP, _ = ants.NewPool(10)
}

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func main() {

	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	if err != nil {
		return
	}
	defer p.Release()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}

	wg.Wait()
	fmt.Println(sum)

}
