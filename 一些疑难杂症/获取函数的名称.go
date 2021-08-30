package main

import (
	"fmt"
	"github.com/golang/glog"
	"runtime"
	"strings"
)

func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	split := strings.Split(f.Name(), ".")
	if len(split) > 1 {
		return split[1]
	}

	return f.Name()
}

func testName() {
	testNext()

	name := RunFuncName()
	fmt.Println("this func name is ", name)
	fmt.Println("this is my test name....")
}

//再次嵌套的函数
func testNext() {
	name := RunFuncName()
	fmt.Println("this func name is ", name)
}

func main() {

	testName()

}

func CatchPanic() {
	if err := recover(); err != nil {
		if pc, file, line, ok := runtime.Caller(2); ok {
			pcName := runtime.FuncForPC(pc).Name()

			glog.Errorf(fmt.Sprintf("catch panic error: %v\nFILE: %s\nLINE: %d\nFUNCNAME: %s", err, file, line, pcName))
		}
	}
}
