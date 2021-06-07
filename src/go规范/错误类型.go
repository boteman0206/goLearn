package main

import (
	"fmt"
)

type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func open(file string) error {
	return errNotFound{file: file}
}

func use() {
	if err := open("testfile.txt"); err != nil {
		if _, ok := err.(errNotFound); ok {
			// handle
			fmt.Println("=======this is err========", err.Error())
		} else {
			panic("unknown error")
		}
	}
}

func main() {
	use()
}
