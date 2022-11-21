package main

import (
	"fmt"
	"os"
)

func severityFrom(args []string) string {
	var s string
	fmt.Println("args: ", args)
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}

func main() {
	from := severityFrom(os.Args)
	fmt.Println(from)
}
