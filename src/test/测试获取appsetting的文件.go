package main

import (
	"fmt"
	"github.com/limitedlee/microservice/common/config"
)

func main() {
	baseURL, err := config.Get("BaseUrl")

	fmt.Println("error : ", err, " baseUrl : ", baseURL)
}
