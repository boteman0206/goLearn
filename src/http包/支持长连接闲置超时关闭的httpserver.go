package main

import (
	"log"
	"net/http"
	"time"
)

/**

显然上面的server处理方式“太过霸道”，对于想要复用连接，提高请求和应答传输效率的client而言，上面的“一刀切”机制并不合理。
那么是否有一种机制可以让http server即可以对高密度传输数据的连接保持keep-alive，又可以及时清理掉那些长时间没有数据传输的idle连接，
释放占用的系统资源呢？我们来看下面这个go实现的server：




// IdleTimeout是当启用keep-alive时等待下一个请求的最大时间。
// 如果IdleTimeout为零，则使用ReadTimeout的值。如果两者都是
// 零，则没有超时。
IdleTimeout time.Duration

*/

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a request from:", r.RemoteAddr, r.Header)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", Index)
	var s = http.Server{
		Addr:        ":8080",
		Handler:     http.HandlerFunc(Index),
		IdleTimeout: 5 * time.Second,
	}
	s.ListenAndServe()
}
