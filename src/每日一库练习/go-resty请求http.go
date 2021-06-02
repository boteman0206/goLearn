package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

/**
todo https://www.cnblogs.com/jiujuan/p/14583605.html#1015161722
https://github.com/go-resty/resty#various-post-method-combinations
go-resty 有很多特性：

	发起 GET, POST, PUT, DELETE, HEAD, PATCH, OPTIONS, etc. 请求
	简单的链式书写
	自动解析 JSON 和 XML 类型的文档
	上传文件
	重试功能
	客户端测试功能
	Resty client
	Custom Root Certificates and Client Certificates
*/

func main() {

	client := resty.New() // 创建一个restry客户端
	resp, err := client.R().EnableTrace().Get("https://httpbin.org/get")

	fmt.Println("resp : ", resp, " error : ", err)
	fmt.Println(" proto : ", resp.Proto(), resp.Body(), resp.Status())

	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)

	StrongGet()

	MyPost()

}

//todo 增强的get
func StrongGet() {
	client := resty.New()
	//client := resty.NewWithClient()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit":   "20",
			"sort":    "name",
			"order":   "asc",
			"random":  strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/search_result")
	fmt.Println(" resp : ", resp, " error : ", err)
	// Request.SetQueryString method
	resp1, err1 := client.R().
		SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/show_product")

	fmt.Println("resp 1 : ", resp1, " error 1 : ", err1)
	// 解析返回的内容，内容是json解析到struct
	//resp2, err2 := client.R().
	//	SetResult(result).
	//	ForceContentType("application/json").
	//	Get("v2/alpine/mainfestes/latest")
}

func MyPost() {
	client := resty.New()

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"testuser", "password":"testpass"}`).
		//SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		Post("https://myapp.com/login")
	fmt.Println("post data : ", resp, err)
}
