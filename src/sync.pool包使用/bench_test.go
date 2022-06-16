package main

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}
var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

/**
go  test -v -bench="." .\bench_test.go




// 查看使用的内存情况
go  test -v -benchmem  -bench="."   .\bench_test.go
			                               执行次数           执行时间                内存分配
BenchmarkUnmarshal-16                      13616             88722 ns/op            1400 B/op          8 allocs/op
BenchmarkUnmarshalWithPool-16              13519             88540 ns/op             248 B/op          7 allocs/op

因为 Student 结构体内存占用较小，内存分配几乎不耗时间。而标准库 json 反序列化时利用了反射，效率是比较低的，占据了大部分时间，因此两种方式最终的执行时间几乎没什么变化。
但是内存占用差了一个数量级，使用了 sync.Pool 后，内存占用仅为未使用的 234/5096 = 1/22，对 GC 的影响就很大了。

*/
