package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// 接受参数的os.Args包
	strings := os.Args
	fmt.Println("main ===========run ")
	fmt.Println("strings : ", strings)

	/**


	E:\goProject\goLearn\src\flag获取命令行的参数>go build demo.go
	E:\goProject\goLearn\src\flag获取命令行的参数>demo.exe abc num
	main ===========run
	strings :  [demo.exe abc num]

	*/

	//使用flag包来接受参数
	var user string
	var pwd string

	flag.StringVar(&user, "user", "", "获取用户名")
	flag.StringVar(&pwd, "pwd", "默认值", "获取密码")

	flag.Parse()

	fmt.Println("user : ", user, " pwd: ", pwd)

	/**

		 E:\goProject\goLearn\src\flag获取命令行的参数>demo.exe -u "jack" -pwd "qeweq
	e"
		main ===========run
		strings :  [demo.exe -u jack -pwd qeweqe]
		user :  jack  pwd:  qeweqe

	*/
}
