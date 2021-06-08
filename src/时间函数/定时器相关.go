package main

import (
	"fmt"
	"time"
)

//func main() {
//	timer := time.NewTimer(time.Second * 5)
//	for {
//		//t.Reset(time.Second * 5)
//_
//		select {
//		//case <- someDone:
//			// do something
//		case <-timer.C:
//			return
//		default:
//			fmt.Println("default run ")
//		}
//
//	}
//
//
//
//}

//func main() {
//	for {
//		//  todo 错误的 一些操作 ，会产生数以万计的timer
//		timeout := time.After(30 * time.Second)
//		select {
//		//case <- someDone:
//			// do something
//		case <-timeout:
//			return
//		}
//	}
//}

//
//func main() {
//	timer1 := time.NewTimer(2 * time.Second)
//	go func() {
//		timer1.Stop()  // todo 产生死锁， 因为程序就会一直死锁了，因为 timer1.Stop 并不会关闭 channel C，使程序一直阻塞在 timer1.C 上。
//	}()
//	<-timer1.C
//
//	println("done")
//}

//todo Stop 的正确的使用方式：
func main() {
	timer1 := time.NewTimer(2 * time.Second)
	go func() {
		if !timer1.Stop() {
			<-timer1.C
		}
	}()

	select {
	case <-timer1.C:
		fmt.Println("expired")
	default:
	}
	println("done")
}
