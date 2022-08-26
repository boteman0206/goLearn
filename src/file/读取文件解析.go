package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"strings"
)

type CategoryIds struct {
	Id  int `json:"id"`
	Pid int `json:"pid"`
}

type PareCategoryIds struct {
	Id    int           `json:"id"`
	Pid   int           `json:"pid"`
	Child []CategoryIds `json:"child"`
}

func main() {

	file, err := ioutil.ReadFile("D:\\RpPet\\gitProject\\goLearn\\src\\file\\json.txt")
	if err != nil {
		return
	}

	//fmt.Println(file)

	var data []PareCategoryIds
	err = json.Unmarshal(file, &data)
	if err != nil {
		return
	}

	var ids []int
	for i := range data {
		fmt.Println(i, " 列数： ", data[i])
		categoryIds := data[i]

		ids = append(ids, categoryIds.Id)
		ids = append(ids, categoryIds.Pid)

		for i2 := range categoryIds.Child {
			c := categoryIds.Child[i2]
			ids = append(ids, c.Id)
			ids = append(ids, c.Pid)
		}
	}
	fmt.Println("ids: ", len(ids), "  ", ids)

	// 去除一下重复
	var mapIds = make(map[int]struct{}, len(ids))
	for i := range ids {
		mapIds[ids[i]] = struct{}{}
	}
	fmt.Println(len(mapIds))

	var lastIds []string
	for k, _ := range mapIds {
		//fmt.Println(k)
		lastIds = append(lastIds, cast.ToString(k))
	}
	fmt.Println("最后的ids：", strings.Join(lastIds, ","))

	//	24406,24467,24471,24400,24460,24462,24716,24396,24427,24401,24415,24428,24434,24449,24463,22049,24397,24454,24438,24446,24429,24475,24421,24426,24424,23412,24402,24444,24455,24457,24472,23414
	//	,24476,24412,24436,24473,24399,24405,24461,24408,24437,24452,24459,23410,24431,24465,24414,24430,24416,24442,24445,24447,24466,24477,24403,24453,24468,24474,24718,24478,24479,24448,24450,24464,24411,2442
	//	5,24404,24433,24407,24418,24451,24420,24469,24717,24422,24417,24419,24458,24470,24435,24443,24441,24432,24439,24440,23415,24409,24398,24423,24410,24719,24413,24456
	////
}
