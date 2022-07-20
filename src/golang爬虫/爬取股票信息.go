package main

import (
	"fmt"
	"github.com/asmcos/requests"
	"regexp"
)

var (
	url = "http://www.iwencai.com/stockpick/search"
)

func main() {

	req := requests.Requests()
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Set("Cookie", "ta_random_userid=dh3uedx0g1; WafStatus=0; PHPSESSID=2ed2dbee31f3f15156255ed8262a5e53; cid=2f84fee9963bdbdd2c07d69481374d381658196345; ComputerID=2f84fee9963bdbdd2c07d69481374d381658196345; v=A6fh2kiDiWjROA3kfF3YWz4RNtB0LH561QH_gnkUweRdI8mOgfwLXuXQj52K")
	req.Header.Set("Referer", "http://www.iwencai.com/stockpick/search?typed=0&preParams=&ts=1&f=1&qs=result_original&selfsectsn=&querytype=stock&searchfilter=&tid=stockpick&w=%E9%95%BF%E5%AE%89%E6%B1%BD%E8%BD%A6")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Host", "www.iwencai.com")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	//
	//p := requests.Params{
	//	"typed":        "0",
	//	"preParams":    "",
	//	"ts":           "1",
	//	"f":            "1",
	//	"qs":           "result_original",
	//	"selfsectsn":   "",
	//	"querytype":    "stock",
	//	"searchfilter": "",
	//	"tid":          "stockpick",
	//	"w":            "长安汽车",
	//}

	get, err := req.Get("http://stockpage.10jqka.com.cn/000625/")
	if err != nil {
		return
	}

	fmt.Println(get.Text())
	fmt.Println(get.SaveFile("./a.txt"))

	regexp.Compile("")
}
