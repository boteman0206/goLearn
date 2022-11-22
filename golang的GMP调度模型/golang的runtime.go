package main

/**

Go Runtime 的总览
	golang 的 runtime 在 golang 中的地位类似于 Java 的虚拟机，不过 go runtime 不是虚拟机. golang 程序生成可执行文件在指定平台上即可运行，效率很高， 它和 c/c++ 一样编译出来的是二进制可执行文件. 我们知道运行 golang 的程序并不需要主机安装有类似 Java 虚拟机之类的东西，那是因为在编译时，golang 会将 runtime 部分代码链接进去.


golang 的 runtime 核心功能包括以下内容:
	1: 协程(goroutine)调度(并发调度模型)
	2: 垃圾回收(GC)
	3: 内存分配
	4: 使得 golang 可以支持如 pprof、trace、race 的检测
	5: 支持 golang 的内置类型 channel、map、slice、string等的实现
	 等等

*/
