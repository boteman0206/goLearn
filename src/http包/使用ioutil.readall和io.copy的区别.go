package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//https://cloud.tencent.com/developer/article/1929867
/**
ioutil.ReadAll源码：
1：创建一个 512 字节的 buf；
2：不断读取内容到 buf，当 buf 满的时候，会追加一个元素，促使其重新分配内存；
3：直到结尾或报错，则返回；
知道了执行步骤，但想要分析其性能问题，还需要了解 Go 切片的扩容策略，如下：

如果期望容量大于当前容量的两倍就会使用期望容量；
如果当前切片的长度小于 1024 就会将容量翻倍；
如果当前切片的长度大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；
也就是说，如果待拷贝数据的容量小于 512 字节的话，性能不受影响。但如果超过 512 字节，就会开始切片扩容。数据量越大，扩容越频繁，性能受影响越大。

如果数据量足够大的话，内存可能就直接撑爆了，这样的话影响就大了。
那有更好的替换方案吗？当然是有的，我们接着往下看。
io.Copy
可以使用 io.Copy 函数来代替，源码定义如下：
如果源实现了 WriteTo 方法，则直接调用 WriteTo 方法；
同样的，如果目标实现了 ReaderFrom 方法，则直接调用 ReaderFrom 方法；
如果 buf 为空，则创建 32KB 的 buf；
最后就是循环 Read 和 Write；


*/

func main() {
	res, err := http.Get("http://www.baidu.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	robots, err := io.ReadAll(res.Body)
	//robots, err := io.Copy(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
