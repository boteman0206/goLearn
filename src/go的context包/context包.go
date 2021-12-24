package main

import (
	"context"
	"fmt"
	"time"
)

/**
type Context interface {

	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}

	Err() error

	Value(key interface{}) interface{}
}

1: Deadline方法是获取设置的截止时间的意思，第一个返回式是截止时间，到了这个时间点，Context会自动发起取消请求；第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。
2: Done方法返回一个只读的chan，类型为struct{}，我们在goroutine中，如果该方法返回的chan可以读取，则意味着parent context已经发起了取消请求，我们通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。之后，Err 方法会返回一个错误，告知为什么 Context 被取消。
3: Err方法返回取消的错误原因，因为什么Context被取消。
4: Value方法获取该Context上绑定的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。

四大函数
	1：func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
		WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。
	2：func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
		WithDeadline函数，和WithCancel差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消Context 当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。
	3：func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
		WithTimeout和WithDeadline基本上一样，这个表示是超时自动取消，是多少时间后自动取消Context的意思。
	4：func WithValue(parent Context, key, val interface{}) Context
		WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context，这个绑定的数据可以通过Context.Value方法访问到，这是我们实际用经常要用到的技巧，一般我们想要通过上下文来传递数据时，可以通过这个方法，如我们需要tarce追踪系统调用栈的时候。





*/

var key = "ctx-key"

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case ctxData := <-ctx.Done():
			//get value
			fmt.Println(ctx.Value(key), "is cancel", " ctxData: ", ctxData)

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}
