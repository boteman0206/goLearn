package main

import (
	"fmt"
	"github.com/spf13/cast"
	"godemo/src/xorm测试/models"
	"sync"
)

func main() {

	s1 := models.WebsocketConns

	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &models.WebsocketConns)

	s2 := models.MapData
	fmt.Printf("%p\n", &s2)
	fmt.Printf("%p\n", &models.MapData)

	//fmt.Println(s1.Load("name"))
	////fmt.Println(s1.Load("name1"))
	//group := new(sync.WaitGroup)
	//group.Add(2000000)
	//go func() {
	//	for i:=0; i<1000000; i++  {
	//		group.Done()
	//		fmt.Println(s1.Load(i))
	//	}
	//}()
	//
	//go func() {
	//	for i:=0; i<1000000; i++  {
	//		group.Done()
	//		models.WebsocketConns.Store(i,  "hello " + cast.ToString(i))
	//	}
	//}()
	//
	//group.Wait()

	lock := sync.RWMutex{}

	group := new(sync.WaitGroup)
	group.Add(2000000)

	go func() {
		for i := 0; i < 1000000; i++ {
			group.Done()
			lock.Lock()
			//models.WebsocketConns.Store(i,  "hello " + cast.ToString(i))
			models.MapData[i] = "hello " + cast.ToString(i)
			lock.Unlock()
		}
	}()

	go func() {
		lock.RLock()
		defer lock.RUnlock()
		for i := 0; i < 1000000; i++ {
			group.Done()

			fmt.Println("---- ", i, s2[i])

		}
	}()

	group.Wait()
}
