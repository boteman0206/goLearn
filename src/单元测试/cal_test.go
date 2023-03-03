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
	//t.SkipNow()  // todo 表示测试的时候跳过当前的方法不执行 执行 go test -v 的时候可以看到当前方法没有执行
	AddNum(1, 10)
	t.Logf("test add num ....")
}

/*
*

	顺序指向的方法

// 这个方式保证测试的顺序, 比如某个test测试依赖于另一个test的结果
*/
func TestShunxun(t *testing.T) {
	t.Run("a1", func(t *testing.T) { fmt.Println("1") })
	t.Run("a2", func(t *testing.T) { fmt.Println("2") })
	t.Run("a3", func(t *testing.T) { fmt.Println("3") })
}

/**
todo 如果要先执行某个test，则用 TestMain(m *testing.M) ，
 通常执行 文件打开 or 登录 or 数据库连接 等等
*/

func TestMain(m *testing.M) {
	fmt.Println("test main first") // 会在文件的开始执行
	m.Run()                        //不加 m.Run() 则不能进行后面的test
}

/*
*
todo 测试性能是有 *test.B
【BenchMark】

	把一个BenchMark重复很多次（不定，程序觉得跑到稳态了就算完，所以BenchMark里面最好是确定的执行时间，不然永远跑不到稳态，永远也跑不完），用来测试性能
*/
func BenchmarkAll(b *testing.B) { //函数名BenchmarkXXX，B大写，XXX首字母大写
	for n := 0; n < b.N; n++ {
		Print1to20()
	}
}
func Print1to20() int {
	res := 0
	for i := 1; i <= 20; i++ {
		res += i
	}
	return res
}

//todo 命令行执行 go test -bench=. -benchmem  可以查看程序的执行时间和内存分配   （程序只会跑带BenchMark的部分）

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
	go test -v -run  cal_test.go TestAddNum
6: 形式参数必须是t *testing.T

*/

// todo 单元测试可以配合 https://github.com/stretchr/testify   asset来断言实现这个包实现了断言

func TestDemo(t *testing.T) {

	t.Log("开始。。。")
	t.Skip("直接跳出了。。。")
	t.FailNow()
	t.Log("结束。。。")
}

// todo 测试覆盖率并输出文件内容到fib.out go test -covermode=count -coverprofile fib.out
