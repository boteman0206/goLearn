package main

import "io"

type myWriter struct {
}

func (w myWriter) Write(p []byte) (n int, err error) {
	return
}

func main() {

	//总结一下，可通过在代码中添加类似如下的代码，用来检测类型是否实现了接口：

	// 检查 *myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = (*myWriter)(nil)

	// 检查 myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = myWriter{}

}
