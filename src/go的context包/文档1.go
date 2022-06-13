package main

import (
	"context"
	"fmt"
	"time"
)

type CompareDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

func main() {

	background := context.Background()
	//todo1 := context.TODO()

	dto := CompareDto{
		Name: "jack",
		Age:  12,
	}

	// valueCtx
	value1 := context.WithValue(background, "v1", dto)
	value := context.WithValue(value1, "v2", "多次存储") // todo 多次存储数据

	fmt.Println("value: ", value.Value("v1"), " v2: ", value.Value("v2"))

	//cancelCtx
	cancel, cancelFunc := context.WithCancel(value)
	fmt.Println("cancel value is : ", cancel.Value("v1"), cancelFunc)
	go func() {
		for {
			select {
			case <-cancel.Done():
				fmt.Println(" 取消了 -----------------")
				return
			default:
				fmt.Println("没有取消")
			}
		}
	}()
	done := cancel.Done()
	fmt.Println("done: ", done, " done == closedchan ", cancel.Done() == closedchan, " err: ", cancel.Err(), " 不会取消的ctx： ", background.Done())
	time.Sleep(1 * time.Microsecond)
	cancelFunc() // 取消之后
	fmt.Println("取消之后 done: ", done, "done == closedchan", cancel.Done() == closedchan, " err: ", cancel.Err(), " 不会取消的ctx： ", background.Done())

	time.Sleep(10 * time.Second)
}

/**
1： type Context interface 是一个接口，里面定义了四个方法
	Deadline() (deadline time.Time, ok bool)     // 返回 context 是否设置了超时时间以及超时的时间点 // 如果没有设置超时，那么 ok 的值返回 false // 每次调用都会返回相同的结果
	Done() <-chan struct{} // 如果 context 被取消，这里会返回一个被关闭的 channel 如果是一个不会被取消的 context，那么这里会返回 nil 每次调用都会返回相同的结果
	Err() error // 返回 done() 的原因   // 如果 Done() 对应的通道还没有关闭，这里返回 nil // 如果通道关闭了，这里会返回一个非 nil 的值：// - 若果是被取消掉的，那么这里返回 Canceled 错误// - 如果是超时了，那么这里返回 DeadlineExceeded 错误 // 一旦被赋予了一个非 nil 的值之后，每次调用都会返回相同的结果
	Value(key interface{}) interface{}   // 获取 context 中保存的 key 对应的 value，如果不存在则返回 nil // 每次调用都会返回相同的结果

2： emptyCtx：实现了Context方法返会的都是nil  emptyCtx 不是一个结构体，它只是 int 类型的一个别名，实现的 Context 的四个方法都是返回 nil 或者默认值：

3：background和todo：都是emptyCtx
var (
    background = new(emptyCtx) // 通常可以用于 main 函数、初始化和测试，作为请求上下文的最顶层（根节点）。
    todo       = new(emptyCtx) // 当你不知道需要传入什么样的 context 的时候，就可以使用它，它可以随时被替换成其他类型的 context。
)


4: valueCtx 相比 emptyCtx，valueCtx 要稍微复杂一些，它维护了一个键值对，可以保存一组 kv（只有一组），其结构体类型定义如下：
结构：
type valueCtx struct {
    // 父 ctx
    Context
    // kv
    key, val interface{}
}
创建 valueCtx 的函数如下：
WithValue(Context, interface{}, interface{}) // value := context.WithValue(background, "name", "lucy")
一个 ctx 只能保存一对 kv，那么如果我们想要在 ctx 中保存多个 kv 键值对该怎么办？
value1 := context.WithValue(background, "v1", "第一次")
value2 := context.WithValue(value1, "v2", "第二次") // todo 多次存储数据，因为查找的时候会向上查找父类的key，value


5： cancelCtx
在讲 cancelCtx 之前，我们先看一下 canceler 接口，因为 cancelCtx 正是实现了这个接口：
type canceler interface {
    cancel(removeFromParent bool, err error) // 取消操作，第一个参数代表取消的时候是否要将当前 ctx 从父 ctx 维护的子 ctx 中移除，第二个参数代表要传给 ctx 的 err（通过 Context.Err() 方法可以捕获）；
    Done() <-chan struct{}  // 与 Context.Done()  一致
}
构建方式： 通过 WithCancel 函数我们可以返回一个 cancelCtx 的实例：
context.WithCancel(background)
type cancelCtx struct {
    Context // 父 ctx
    mu       sync.Mutex   // 通过互斥锁来保证对下面三个 field 操作的安全性
    done     atomic.Value    // 保存一个 chan struct{}，第一个 cancel() 调用会关闭这个通道
    children map[canceler]struct   // 维护所有子 canceler // 当前 ctx 被取消之后，它的所有子 canceler 都会被取消，并且这个属性会被置空
    err      error   // 第一个 cancel() 调用会赋值
}



*/
