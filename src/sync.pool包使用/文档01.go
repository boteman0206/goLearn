package main

/**
1 sync.Pool 的使用场景
一句话总结：保存和复用临时对象，减少内存分配，降低 GC 压力。
是 sync 包下的一个组件，可以作为保存临时取还对象的一个“池子”。个人觉得它的名字有一定的误导性，因为 Pool 里装的对象可以被无通知地被回收，可能 sync.Cache 是一个更合适的名字。

sync.Pool 是一个内存池。通常内存池是用来防止内存泄露的（例如C/C++)。
sync.Pool 这个内存池却不是干这个的，带 GC 功能的语言都存在垃圾回收 STW 问题，需要回收的内存块越多，STW 持续时间就越长。
如果能让 new 出来的变量，一直不被回收，得到重复利用，是不是就减轻了 GC 的压力。

正确的使用示例（下面的demo选自gin）
`func(engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){`
 `c:=engine.pool.Get().(*Context)`
 `c.writermem.reset(w)`
 `c.Request=req`
 `c.reset()`
 `engine.handleHTTPRequest(c)`
 `engine.pool.Put(c)`
`}`



使用场景
上面说到 sync.Pool 业务开发中不是一个常用结构，我们业务开发中没必要假想某块代码会有强烈的性能问题，一上来就用 sync.Pool 硬怼。
sync.Pool 主要是为了解决 Go GC 压力过大问题的，所以一般情况下，当线上高并发业务出现 GC 问题需要被优化时，才需要用 sync.Pool 出场。
1: 当多个 goroutine 都需要创建同⼀个对象的时候，如果 goroutine 数过多，导致对象的创建数⽬剧增，进⽽导致 GC 压⼒增大。形成 “并发⼤－占⽤内存⼤－GC 缓慢－处理并发能⼒降低－并发更⼤”这样的恶性循环。
2: 在这个时候，需要有⼀个对象池，每个 goroutine 不再⾃⼰单独创建对象，⽽是从对象池中获取出⼀个对象（如果池中已经有的话）。

使用注意点
sync.Pool 同样不能被复制。
好的使用习惯，从 pool.Get 出来的值进行数据的清空（reset），防止垃圾数据污染。


*/

func main() {

}
