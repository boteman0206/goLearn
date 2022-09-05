package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name string `json:"name,omitempty"`       // todo  omitempty 标签会在值为默认值的时候直接忽略序列化
	Age  int32  `json:"age,string,omitempty"` //todo 虽然age为int，但是解析序列化之后为string
}

func (a *Author) MarshalJSON() ([]byte, error) {
	return []byte{1, 2, 3}, nil
}

func main() {

	author := Author{
		"",
		15,
	}

	marshal, err := json.Marshal(author)
	if err != nil {
		return
	}

	fmt.Println(string(marshal))

}

/**

参考文档; https://draveness.me/golang/docs/part4-advanced/ch09-stdlib/golang-json/

1： JSON 标准库中提供了 encoding/json.Marshaler 和 encoding/json.Unmarshaler 两个接口分别可以影响 JSON 的序列化和反序列化结果：
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
在 JSON 序列化和反序列化的过程中，它会使用反射判断结构体类型是否实现了上述接口，如果实现了上述接口就会优先使用对应的方法进行编码和解码操作，除了这两个方法之外，
Go 语言其实还提供了另外两个用于控制编解码结果的方法，即 encoding.TextMarshaler 和 encoding.TextUnmarshaler：
type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}

type TextUnmarshaler interface {
	UnmarshalText(text []byte) error
}
一旦发现 JSON 相关的序列化方法没有被实现，上述两个方法会作为候选方法被 JSON 标准库调用并参与编解码的过程。
总的来说，我们可以在任意类型上实现上述这四个方法自定义最终的结果，后面的两个方法的适用范围更广，但是不会被 JSON 标准库优先调用。



2：标准库会使用如下所示的 encoding/json.parseTag 来解析标签：
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

从该方法的实现中，我们能分析出 JSON 标准库中的合法标签是什么形式的：标签名和标签选项都以 , 连接，最前面的字符串为标签名，后面的都是标签选项。



*/
