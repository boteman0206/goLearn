package bufio使用

/**

1.io库属于底层接口定义库，其作用是是定义一些基本接口和一些基本常量，并对这些接口的作用给出说明，常见的接口有Reader、Writer等。一般用这个库只是为了调用它的一些常量，比如io.EOF。

2.ioutil库包含在io目录下，它的主要作用是作为一个工具包，里面有一些比较实用的函数，
比如 ReadAll(从某个源读取数据)、ReadFile（读取文件内容）、WriteFile（将数据写入文件）、ReadDir（获取目录）

3.os库主要是跟操作系统打交道，所以文件操作基本都会跟os库挂钩，比如创建文件、打开一个文件等。这个库往往会和ioutil库、bufio库等配合使用

4.bufio库可以理解为在io库上再封装一层，加上了缓存功能。它可能会和ioutil库和bytes.Buffer搞混。

4.1 bufio VS ioutil库：两者都提供了对文件的读写功能，唯一的不同就是bufio多了一层缓存的功能，这个优势主要体现读取大文件的时候
（ioutil.ReadFile是一次性将内容加载到内存，如果内容过大，很容易爆内存）

4.2 bufio VS bytes.Buffer：两者都提供一层缓存功能，它们的不同主要在于 bufio 针对的是文件到内存的缓存，
而 bytes.Buffer 的针对的是内存到内存的缓存（个人感觉有点像channel，你也可以发现 bytes.Buffer 并没有提供接口将数据写到文件）。


5.bytes和strings库：这两个库有点迷，首先它们都实现了Reader接口，所以它们的不同主要在于针对的对象不同，bytes针对的是字节，strings针对的是字符串（它们的方法实现原理很相似）。
另一个区别就是 bytes还带有Buffer的功能，但是 strings没提供。
注：关于Reader和Writer接口，可以简单理解为读取源和写入源，即只要实现Reader里面的Read方法，这个东西就可以作为一个读取源，里面可以包含数据并被我们读取；Writer亦是如此。





*/
