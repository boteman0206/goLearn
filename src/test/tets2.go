package main

import (
	//"database/sql"
	"fmt"
	//"godemo/src/model"
	"strconv"
	//"time"

	kit "github.com/tricobbler/rp-kit"
)

func main() {
	defer func() {
		fmt.Println("recovered:")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("not good")
}
