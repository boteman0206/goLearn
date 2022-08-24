package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return
	}

	fmt.Println(newUUID)
}
