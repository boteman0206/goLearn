package main

/**
https://tonybai.com/2021/04/02/go-http-client-connection-control/ 文档

最初早期的http 1.0协议只支持短连接，即客户端每发送一个请求，就要和服务器端建立一个新TCP连接，请求处理完毕后，该连接将被拆除。
显然每次tcp连接握手和拆除都将带来较大损耗，为了能充分利用已建立的连接，后来的http 1.0更新版和http 1.1支持在http请求头中加入Connection: keep-alive来告诉对方这个请求响应完成后不要关闭链接，下一次还要复用这个连接以继续传输后续请求和响应。
后HTTP协议规范明确规定了HTTP/1.0版本如果想要保持长连接，需要在请求头中加上Connection: keep-alive，而HTTP/1.1版本将支持keep-alive长连接作为默认选项，有没有这个请求头都可以。

1： http包默认启用keep-alive

2. http client端基于非keep-alive连接发送请求
client := http.Client{Transport: &http.Transport{DisableKeepAlives: true}}

3. 支持长连接闲置超时关闭的http server

4: client 类型
Client 类型代表 HTTP 客户端。它的零值（ DefaultClient ）是一个可用的使用 DefaultTransport 的客户端。
Client 的 Transport 字段一般会含有内部状态（缓存 TCP 连接），todo 因此 Client 类型值应尽量被重用而不是每次需要都创建新的。Client 类型值可以安全的被多个 go 程同时使用。
Client 类型的层次比 RoundTripper 接口（如 Transport ）高，还会管理 HTTP 的 cookie 和重定向等细节。

5: 主要结构
type Client struct {
    // Transport 指定执行独立、单次 HTTP 请求的机制。
    // 如果 Transport 为 nil，则使用 DefaultTransport 。
    Transport RoundTripper

    // CheckRedirect 指定处理重定向的策略。
    // 如果 CheckRedirect 不为 nil，客户端会在执行重定向之前调用本函数字段。
    // 参数 req 和 via 是将要执行的请求和已经执行的请求（切片，越新的请求越靠后）。
    // 如果 CheckRedirect 返回一个错误，本类型的 Get 方法不会发送请求 req，
    // 而是返回之前得到的最后一个回复和该错误。（包装进 url.Error 类型里）
    //
    // 如果CheckRedirect为nil，会采用默认策略：连续10此请求后停止。
    CheckRedirect func(req *Request, via []*Request) error

    // Jar 指定 cookie 管理器。
    // 如果Jar为nil，请求中不会发送 cookie ，回复中的 cookie 会被忽略。
    Jar CookieJar

    // Timeout 指定本类型的值执行请求的时间限制。
    // 该超时限制包括连接时间、重定向和读取回复主体的时间。
    // 计时器会在 Head 、 Get 、 Post 或 Do 方法返回后继续运作并在超时后中断回复主体的读取。
    //
    // Timeout 为零值表示不设置超时。
    //
    // Client 实例的 Transport 字段必须支持 CancelRequest 方法，
    // 否则 Client 会在试图用 Head 、 Get 、 Post 或 Do 方法执行请求时返回错误。
    // 本类型的 Transport 字段默认值（ DefaultTransport ）支持 CancelRequest 方法。
    Timeout time.Duration
}

6：Do 方法 func (c *Client) Do(req *Request) (resp *Response, err error)
调用者应该在读取完 resp.Body 后关闭它。如果返回值 resp 的主体未关闭，c 下层的 RoundTripper 接口（一般为 Transport 类型）可能无法重用 resp 主体下层保持的 TCP 连接去执行之后的请求。


7:优化使用http读取数据
https://blog.thinkeridea.com/201901/go/you_ya_de_du_qu_http_qing_qiu_huo_xiang_ying_de_shu_ju.html

*/
func main() {
	//client := http.Client{Transport: &http.Transport{DisableKeepAlives: true}}

	//http.NewRequest()
}
