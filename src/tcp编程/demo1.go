package main

import (
	"fmt"
	"net"
)

func processRead(conn net.Conn) {
	defer conn.Close()

	for {
		//服务器等待客户端
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[:n]))

	}

}

func main() {
	fmt.Println("服务器开始监听。。。")
	listener, e := net.Listen("tcp", "127.0.0.1:9091")
	fmt.Println(listener, "e:", e)

	// 延时关闭listen
	defer listener.Close()

	for {
		conn, i := listener.Accept() // 等待客户端连接

		fmt.Println("conn : ", conn, " i:", i)
		fmt.Println(conn.RemoteAddr(), "连接。。。。")
		go processRead(conn)
	}

}
