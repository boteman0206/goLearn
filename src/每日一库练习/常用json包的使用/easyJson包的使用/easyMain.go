package main

import (
	"fmt"
	"godemo/testDemo"
	"time"
)

/**
go get -u github.com/mailru/easyjson/
-lower_camel_case:将结构体字段field首字母改为小写。如Name=>name。
-build_tags string:将指定的string生成到生成的go文件头部。
-no_std_marshalers：不为结构体生成MarshalJSON/UnmarshalJSON函数。
-omit_empty:没有赋值的field可以不生成到json，否则field为该字段类型的默认值。
-output_filename:定义生成的文件名称。
-pkg:对包内指定有`//easyjson:json`结构体生成对应的easyjson配置。
-snke_case:可以下划线的field如`Name_Student`改为`name_student`。
*/

/**
todo
	一： 必须在结构体上面加上 //easyjson:json
	二： 在结构体的包下面执行  easyjson  -all student.go
	在测试中出错： Bootstrap failed: 20:39: illegal UTF-8 encoding (and 95 more errors) 文件名称不能为中文

*/

func main() {
	s := easyTest.Student{
		Id:   11,
		Name: "qq",
		School: easyTest.School{
			Name: "CUMT",
			Addr: "xz",
		},
		Birthday: time.Now(),
	}
	bt, err := s.MarshalJSON()
	fmt.Println(string(bt), err)
	json := `{"id":11,"s_name":"qq","s_chool":{"name":"CUMT","addr":"xz"},"birthday":"2017-08-04T20:58:07.9894603+08:00"}`
	ss := easyTest.Student{}
	ss.UnmarshalJSON([]byte(json))
	fmt.Println(ss)
}
