package main

import (
	"encoding/json"
	"fmt"
	"github.com/panjf2000/ants"
	"github.com/spf13/cast"
	"time"

	//"github.com/spf13/cast"
	"sync"
	"sync/atomic"
	//"github.com/spf13/cast"
	//"time"
)

type User struct {
	Name string
	Age  int
	Pets []string
}

func main() {

	var wait sync.WaitGroup

	p, _ := ants.NewPool(5)
	defer p.Release()

	chan_user := make(chan User, 10)
	datas := make([]*User, 0)
	var num int32
	MtfailResult := []*User{}

	go func() {
		for true {
			select {
			case data := <-chan_user:
				fmt.Println("data: ", data)
				datas = append(datas, &data)
				//default:
				//	fmt.Println("==========data =")
			}

		}
	}()

	for j := 0; j < 100; j++ {

		fmt.Println(j)
		wait.Add(1)

		user := User{}
		user.Age = j
		_ = p.Submit(func() {

			defer func() {
				wait.Done()
			}()
			user.Name = cast.ToString(user.Age)
			//MtfailResult = append(MtfailResult, &user)
			//time.Sleep( 1* time.Second)
			atomic.AddInt32(&num, 1)
			chan_user <- user
		})

	}

	wait.Wait()

	fmt.Println("启动的线程数： ", p.Running(), len(MtfailResult))
	//time.Sleep(20 * time.Second)
	fmt.Println(StructToString(datas))
	fmt.Println("长度： ", len(datas), num)

}

func StructToString(i interface{}) string {

	bytes, _ := json.Marshal(i)
	return string(bytes)
}

var newPool *ants.Pool
var Num int32

func init() {

	newPool, _ = ants.NewPool(10)
}

func TestFunc() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		i := i
		newPool.Submit(func() {
			defer func() {
				wg.Done()
			}()
			time.Sleep(10 * time.Millisecond)
			fmt.Println(i)

		})

		atomic.AddInt32(&Num, 1)

	}

	wg.Wait()
	fmt.Println("启动的线程数：", newPool.Running(), " ", newPool.Cap())
	fmt.Println("num: ", Num)

}
