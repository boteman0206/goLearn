package main

import (
	"flag"
	"fmt"
)

/*解析命令行参数*/
func main() {
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为空")
	flag.IntVar(&port, "port", 3306, "端口号，默认为空")
	/*转换*/
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)

	// todo 使用 flagDemo.exe -u "jack"  -pwd 123345 -h "127.0.0.1" -port 90 命令启动
}
