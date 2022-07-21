package main

import (
	"bytes"
	"fmt"
	"github.com/RoaringBitmap/roaring"
)

func main() {
	rb1 := roaring.BitmapOf(349, 98, 1, 2, 3, 4, 5, 100, 1000)
	fmt.Println(rb1.String())

	fmt.Println(rb1.Select(1))

	rb := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	buf := new(bytes.Buffer)
	size, err := rb.WriteTo(buf)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("size: ", size, buf.String())

	fmt.Println(string(buf.Bytes()))

}
