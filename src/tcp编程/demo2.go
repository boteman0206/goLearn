package main

import (
	"fmt"
	"net"
)

func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:9091")
	fmt.Println(conn, " e:", e)

	n, err := conn.Write([]byte("hello world!"))

	if err != nil {
		fmt.Println("发送失败")
	}
	fmt.Println(n)

}
