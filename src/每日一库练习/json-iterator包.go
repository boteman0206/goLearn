package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
)

//stringE, err := cast.ToStringMapStringE(sname)
//fmt.Println(stringE, err)

var jsonData = jsoniter.ConfigCompatibleWithStandardLibrary

type ColorGroup struct {
	ID     int `json:"id,string"` // todo 必须指定string类型
	Name   string
	Colors []string
}

func main() {

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err = json_iterator.Marshal(group)
	os.Stdout.Write(b)
	fmt.Println()
	var group1 ColorGroup
	sname := `{"id": "100","name":"jack"}`
	err = json.Unmarshal([]byte(sname), &group1)
	fmt.Println(err, "---", group1)
	fmt.Println(group1)

	err = json_iterator.Unmarshal([]byte(sname), &group1)
	fmt.Println(err, "+++++")
	fmt.Println(group1)

}
