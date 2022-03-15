package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	fmt.Println(buf.String())
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func deepCopy02(dst, src interface{}) error {

	marshal, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(marshal, &dst)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	slice01 := []int{1, 2, 3, 4, 5}
	var slice02 = make([]int, 0)

	deepCopy(&slice02, &slice01) // todo 注意传递指针

	fmt.Println(slice01, " 02: ", slice02)

	copy02 := deepCopy02(&slice02, &slice01) // todo 注意传递指针
	if copy02 != nil {
		fmt.Println(copy02)
	}

	fmt.Println(slice01, " === ", slice02)
}
