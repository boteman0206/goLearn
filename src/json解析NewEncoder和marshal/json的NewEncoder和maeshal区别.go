package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	kit "github.com/tricobbler/rp-kit"
	"io/ioutil"
	"net/http"
	"os"
)

type Use struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//json.Unmarshal进行解码
func HandleUse01(w http.ResponseWriter, r *http.Request) {
	var u Use
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(data, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "姓名：%s，年龄：%d", u.Name, u.Age)

}

//json.NewDecoder解码
func HandleUse02(w http.ResponseWriter, r *http.Request) {
	var u Use
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "姓名：%s，年龄：%d", u.Name, u.Age)

}

/**
区别
	1、json.NewDecoder是从一个流里面直接进行解码，代码精干
	2、json.Unmarshal是从已存在与内存中的json进行解码
	3、相对于解码，json.NewEncoder进行大JSON的编码比json.marshal性能高，因为内部使用pool
场景应用
1、json.NewDecoder用于http连接与socket连接的读取与写入，或者文件读取
2、json.Unmarshal用于直接是byte的输入


Encoder
Encoder 主要负责将结构对象编码成 JSON 数据，我们可以调用 json.NewEncoder(io.Writer) 方法获得一个 Encoder 实例：
func NewEncoder(w io.Writer) *Encoder {
    return &Encoder{w: w, escapeHTML: true}
}
Decoder
Decoder 主要负责将 JSON 数据解析成结构对象，我们可以调用 json.NewDecoder(io.Reader) 方法获得一个 Decoder 实例：
// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
    return &Decoder{r: r}
}


*/
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// 如果Child字段为nil 编码JSON时可忽略
	Child *Person `json:"child,omitempty"`
}

func main() {
	person := Person{
		Name: "John",
		Age:  40,
		Child: &Person{
			Name: "Jack",
			Age:  20,
		},
	}

	// File类型实现了io.Writer接口
	file, _ := os.Create("person.json")

	// 根据io.Writer创建Encoder 然后调用Encode()方法将对象编码成JSON
	json.NewEncoder(file).Encode(&person)
	//todo 上面程序会将结构体对象编码成 JSON 数据，存入 person.json 文件中，程序运行后，会生成下面文件内容：

	var person1 Person

	// File类型也实现了io.Reader接口
	file1, _ := os.Open("person.json")

	// 根据io.Reader创建Decoder 然后调用Decode()方法将JSON解析成对象
	json.NewDecoder(file1).Decode(&person1)

	fmt.Println(person)
	fmt.Println(*person.Child)

	//常用解析
	type user struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	usrd := user{
		Name: "jojo",
		Age:  90,
	}

	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&usrd) // 放到bytes.Buffer中

	fmt.Println("data: ", data.String())

	var user1 user

	json.NewDecoder(&data).Decode(&user1) // 解析出来
	fmt.Println("this is use1 : ", kit.JsonEncode(user1))

}
