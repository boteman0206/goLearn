package main

import (
	"fmt"
	"github.com/golang/glog"
	"runtime"
	"sync"
)

func main() {

	gNum := runtime.NumCPU()
	channelG := make(chan bool, gNum)
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {

		channelG <- true
		wg.Add(1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					glog.Error("DealAsyncTaskListAll-定时任务异常 ")
				}
			}()
			defer func() {
				fmt.Println("释放 close", wg)
				<-channelG
				wg.Done()
			}()
			TaskContent := 11
			switch TaskContent {
			case 1: //todo 商品批量新建 excel导入商品到商品库
				fmt.Println("1   ...")
			case 2: //todo 批量更新
				fmt.Println("2   ...")

			case 11:
				//todo 渠道--商品导出
				//公共使用 operation_file_url存的是财务编码
				fmt.Println("11   ...")
				break
			default:
				//如果以上都没有处理，则更新下最后创建时间
				fmt.Println("default ...")
			}
		}()
	}
	wg.Wait()
	close(channelG)

}
