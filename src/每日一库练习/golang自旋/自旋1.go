package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/**
参考文档： https://ofstack.com/Golang/27085/implementation-of-golang-spin-lock.html

*/
type spinLock uint32

func (sl *spinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}
func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}
func NewSpinLock() sync.Locker {
	var lock spinLock
	return &lock
}

func main() {

	lock := NewSpinLock()

	var num = 0
	var wait sync.WaitGroup
	for i := 0; i < 100; i++ {
		wait.Add(1)
		go func() {
			defer func() {
				wait.Done()
			}()
			lock.Lock()
			defer lock.Unlock()
			num++
			//
			//var buf [64]byte
			//n := runtime.Stack(buf[:], false)
			//idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
			//fmt.Println("idField: ", idField)

		}()
	}

	wait.Wait()

	fmt.Println("自旋锁的num: ", num)

}
