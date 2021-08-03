package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func main() {
	//初始化命令行参数
	flag.Parse()

	//退出时调用，确保日志写入文件中
	defer glog.Flush()

	fmt.Println(glog.V(0))
	glog.V(0).Info("hello, glog")
	if glog.V(0) {
		glog.Info("oooooooooooooooo")
	}
	glog.Warning("warning glog")
	glog.Error("error glog")
	//glog.Fatal("test fatal")
	//
	glog.Infof("info %d", 1)
	glog.Warningf("warning %d", 2)
	glog.Errorf("error %d", 3)

	glog.Info("This is info message")
	glog.Infof("This is info message: %v", 12345)
	glog.InfoDepth(1, "This is info message", 12345)

}

/**

一： go build 编译
二： glog日志.exe -log_dir="./"   //执行exe文件产生log日志

*/
