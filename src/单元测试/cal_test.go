package cal

import (
	"fmt"
	"testing"
)

func TestCalYy(t *testing.T) {

	cal := calYy(10)
	fmt.Println(cal)

	AddNum(1, 2)

	if cal == 10 {
		t.Log("测试实际值！！！")
	}

	t.Logf("popTestCalYy......")

}

func TestAddNum(t *testing.T) {
	AddNum(1, 10)
	t.Logf("test add num ....")
}

/**

E:\goProject\goLearn\src\单元测试>go test -v
=== RUN   TestCalYy
cal run... 10
10
    cal_test.go:14: 测试实际值！！！
--- FAIL: TestCalYy (0.00s)
FAIL
exit status 1
FAIL    _/E_/goProject/goLearn/src/单元测试     0.136s

*/

/**
注意点：
1： 必须要以_test.go结尾
2： 预备测试的文件在同一个包中
3： 函数需要以Test开头
4： 测试单个文件，一定要带上被测试的原文件
	go test -v cal_test.go cal.go
5: 测试单个方法
	go test -v -test.run TestAddNum
6: 形式参数必须是t *testing.T

*/

// todo 单元测试可以配合 https://github.com/stretchr/testify   asset来断言实现这个包实现了断言
