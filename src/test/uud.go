package main

import (
	"fmt"
	"go.uber.org/zap/buffer"
	"sync"
)

//func main() {
//
//
//
//	var str strings.Builder
//
//	str.WriteString("ad")
//	str.WriteString("世界")
//
//	fmt.Println(str.String())
//
//
//	var buf buffer.Pool
//	get := buf.Get()
//	fmt.Println(get)
//
//}

var pool = sync.Pool{
	New: func() interface{} {
		return "1123"
	},
}

func main() {
	t := pool.Get().(string)
	fmt.Println(t)

	//pool.Put("321")
	pool.Put("32112")
	pool.Put("323241")
	pool.Put("3521")

	t2 := pool.Get().(string)
	fmt.Println(t2)

	newPool := buffer.NewPool()

	get := newPool.Get()
	fmt.Println(get)

}
