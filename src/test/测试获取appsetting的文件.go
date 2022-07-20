package main

import (
	"fmt"
	"github.com/limitedlee/microservice/common/config"
)

func main() {
	baseURL, err := config.Get("BaseUrl")

	fmt.Println("error : ", err, " baseUrl : ", baseURL)

	var isContinue int
	if r, err := GetData(); err != nil {
		fmt.Println(isContinue)
	} else if r == 1 {
		isContinue = 1
	}

	fmt.Println(isContinue)
}

func GetData() (int, error) {
	return 1, nil
}
