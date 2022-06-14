package main

/**

原先的map无法支持并发读写

两种方式实现并发读写：
	1： 使用互斥锁包装
	2： 使用sync.Map
sync.Map在读和删除两项性能基准测试上的数据都大幅领先使用sync.Mutex或RWMutex包装的原生map，仅在写入一项上存在一倍的差距。sync.Map是如何实现如此高的读取性能的呢？简单说：空间换时间+读写分离+原子操作(快路径)。

sync.Map的原理简述
	sync.Map底层使用了两个原生map，一个叫read，仅用于读；一个叫dirty，用于在特定情况下存储最新写入的key-value数据:
	1: read(这个map)好比整个sync.Map的一个“高速缓存”，当goroutine从sync.Map中读取数据时，sync.Map会首先查看read这个缓存层是否有用户需要的数据(key是否命中)，如果有(命中)，则通过原子操作将数据读取并返回，这是sync.Map推荐的快路径(fast path)，也是为何上面基准测试结果中读操作性能极高的原因。


*/
