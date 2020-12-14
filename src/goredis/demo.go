package main

import (
	"fmt"
	. "github.com/gomodule/redigo/redis"
)

func main() {

	conn, e := Dial("tcp", "127.0.0.1:6379")
	if e != nil {
		fmt.Println("连接redis失败。。")
		return
	}

	fmt.Println("conn : ", conn)

	// 设置数据
	reply, err := conn.Do("set", "name1", "tom")

	if err != nil {
		fmt.Println("设置失败")
		return
	}
	fmt.Println("reply : ", reply)

	// 获取数据
	s, i := String(conn.Do("get", "name1"))
	fmt.Println(s, i)

}
