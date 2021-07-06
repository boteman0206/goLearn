package main

import (
	"fmt"
	"strings"
)

func Test(name ...string) {

	for i, v := range name {
		fmt.Println(i, " == ", v)
	}

}

func main() {

	Test("jack", "bob", "lucy")

	name := "hello"
	count := strings.Count(name, "") // todo 总的字节数 如果substr为空返回值 +1
	fmt.Println("count : ", count)

	autho := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjA4NzUwMjgsIm1vYmlsZSI6IjE4ODE4MjYxNDYxIiwibmFtZSI6InBlbmd3IiwidXNlcm5vIjoiVV8zRVk3QkNRIn0.SA-7iNS3kcNbmX7FvOIlhFEO_RJfhJzBOL8kE_yomUN3nbmkFFwpeUFhmz-JUFytq7tcdOMG27MX5NYGHUzZ6f-lFYPEEUhKnr5zmpePGSpPGNlJKL6BWQQsRN-v8QcAbupJL6B3vNaJwRu0Pbpc5nJ1HZNLqi5_BSSj7-qoELs"
	index := strings.Index(autho, " ")
	fmt.Println("index : ", index)

	data := autho[strings.Index(autho, " ") : strings.Count(autho, "")-1]
	fmt.Println("data1 :", data)
	data2 := autho[strings.Index(autho, " "):len(autho)]
	fmt.Println("data2 :", data2)

	//TODO 这个包是干嘛用的
	//get, err := config.Get("redis.Addr")
	//fmt.Println(err, " redis get : ", get)

	fmt.Println(autho[0:6])
}
