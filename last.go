package main

import (
	"fmt"
	"godemo/src/localRedis"
	"time"
)

func main() {

	localRedis.SetupDB()

	conn := localRedis.GetRedisConn()
	fmt.Printf("%p", conn)

	//for i:=0; i<100 ; i++  {
	//
	//	conn := localRedis.GetRedisConn()  // 0xc00005c060
	//	fmt.Printf("%p", conn)
	//	fmt.Println("\n")
	//	time.Sleep(10*time.Second)
	//}

	time.Sleep(10 * time.Hour)
}
