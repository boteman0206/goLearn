package main

import (
	"fmt"
	"sync"
)

/**
1: 为什么用 Pool，而不是在运行的时候直接实例化对象呢？
	本质原因：Go 的内存释放是由 runtime 来自动处理的，有 GC 过程。
2: 为什么 sync.Pool 不适合用于像 socket 长连接或数据库连接池?
因为，我们不能对 sync.Pool 中保存的元素做任何假设，以下事情是都可以发生的：
	2.1: Pool 池里的元素随时可能释放掉，释放策略完全由 runtime 内部管理；
	2.2: Get 获取到的元素对象可能是刚创建的，也可能是之前创建好 cache 住的。使用者无法区分；
	2.3" Pool 池里面的元素个数你无法知道；
	所以，只有的你的场景满足以上的假定，才能正确的使用 Pool 。
	sync.Pool 本质用途是增加临时对象的重用率，减少 GC 负担。划重点：临时对象。所以说，像 socket 这种带状态的，长期有效的资源是不适合 Pool 的。
*/
var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
}

func main() {
	initPool()

	/**
	// 申请对象 Get
	Get 方法会返回 Pool 已经存在的对象，如果没有，那么就走慢路径，也就是调用初始化的时候定义的 New 方法（也就是最开始定义的初始化行为）来初始化一个对象。
	*/
	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	/**
	释放对象 Put
	使用对象之后，调用 Put 方法声明把对象放回池子。注意了，这个调用之后仅仅是把这个对象放回池子，池子里面的对象啥时候真正释放外界是不清楚的，是不受外部控制的。
	*/
	pool.Put(p)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person))
}
