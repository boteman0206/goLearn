package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/limitedlee/microservice/common/config"
	kit "github.com/tricobbler/rp-kit"
	"strconv"
	"strings"
	"time"
)

func main() {

	url := fmt.Sprintf("%s%s", config.GetString("bj-scrm-url"), "/api/zlcrm/openapi/vaccinate/get-list")
	BjSignMap(url)

}

func BjSignMap(url string) map[string]string {
	domainUrl := strings.Split(url, "//")[1]
	baseUrl := strings.Split(domainUrl, "/")[0]
	method := strings.Split(url, baseUrl)[1]
	Timestamp := strconv.Itoa(int(time.Now().Unix()))
	//sign := fmt.Sprintf("AppId=%s&Secret=%s&Url=%s&Timestamp=%s&Version=%s", config.GetString("bj.auth.appid"), config.GetString("bj.auth.secret"), method, Timestamp, config.GetString("bj.auth.version"))
	sign := fmt.Sprintf("AppId=%s&Secret=%s&Url=%s&Timestamp=%s&Version=%s",
		"",
		"", method, Timestamp, "1.0.0.0")
	h := md5.New()
	h.Write([]byte(sign))
	md5sign := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	arr := make(map[string]string)
	arr["focus-auth-appid"] = "sz_rppet"
	arr["focus-auth-userid"] = "0"
	arr["focus-auth-username"] = "0"
	arr["focus-auth-version"] = "1.0.0.0"
	arr["focus-auth-url"] = method
	arr["focus-auth-timestamp"] = Timestamp
	arr["focus-auth-sign"] = md5sign
	fmt.Println("arr: ", kit.JsonEncode(arr))
	return arr
}
