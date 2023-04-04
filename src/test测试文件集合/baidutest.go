package main

/**
docker run -d --name kafka -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=127.0.0.1:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://127.0.0.01:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -e TZ="Asia/Shanghai"  -v /etc/localtime:/etc/localtime wurstmeister/kafka:2.12-2.5.0

docker run -d --restart=always --log-driver json-file --log-opt max-size=100m --log-opt max-file=2 --name kafka -p 9002:9002 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=172.18.184.38:2181/kafka -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://172.18.184.38:9002 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9002 -v /etc/localtime:/etc/localtime wurstmeister/kafka:2.12-2.5.0


docker run -d --name kafka --publish 9092:9092 \
--link zookeeper \
--env KAFKA_ZOOKEEPER_CONNECT=172.18.184.38:2181 \
--env KAFKA_ADVERTISED_HOST_NAME=172.18.184.38 \
--env KAFKA_ADVERTISED_PORT=9092  \
--env KAFKA_LOG_DIRS=/kafka/kafka-logs-1 \
-v /usr/local/kafka/logs:/kafka/kafka-logs-1  \
wurstmeister/kafka:2.12-2.5.0




172.18.184.38
*/

import (
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"strings"
	"time"
)

func main() {

	//var CDN_DOMAIN = "https://d1tha5yds2obwc.cloudfront.net/"
	//
	//PrizeUrl := "https://d1tha5yds2obwc.cloudfront.net/a/b/c.txt"
	//if strings.Contains(PrizeUrl, CDN_DOMAIN) {
	//	PrizeUrl = strings.TrimPrefix(PrizeUrl, CDN_DOMAIN)
	//}
	//
	//fmt.Println("------", PrizeUrl)
	//randNum()

	d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//Shuffer(d)

	d = append(d[:2], d[2+1:]...)
	fmt.Println(len(d), "----", d)

	ids := RemoveRepByMap([]string{"1", "2", "3", "1", "4", "5"})
	fmt.Println(ids)

	var relation uint8
	if true {
		relation |= 0b01
	}
	if true {
		relation |= 0b10
	}

	fmt.Println("relation: ", relation)
	fmt.Println("++++", relation&0b10)

}

func RemoveRepByMap(slc []string) []string {
	result := []string{}         //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func randNum() {
	const n = 10000000
	const m = 3000

	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(n)[:m]

	var strs []string
	for i := range nums {
		id := nums[i]
		strs = append(strs, cast.ToString(id))
	}
	join := strings.Join(strs, ",")
	fmt.Println(join)

}

func Shuffer(allList []int) {
	fmt.Println("开始： ", allList)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allList), func(i, j int) {
		allList[j], allList[i] = allList[i], allList[j]
	})
	fmt.Println("结束： ", allList)
}
