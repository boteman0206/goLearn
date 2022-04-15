package main

import (
	"database/sql"
	"fmt"
	"time"
)

func main() {

	type BaseModel struct {
		CreateTime time.Time    `gorm:"comment:'创建时间';type:timestamp;";json:"createTime"`
		UpdateTime time.Time    `gorm:"comment:'修改时间';type:timestamp;";json:"updateTime"`
		BirthTime  sql.NullTime `gorm:"comment:'修改时间';type:timestamp;";json:"updateTime"`
		Remark     string       `gorm:"comment:'备注'";json:"remark"`
	}

	model := BaseModel{}

	//model.CreateTime = &time.Now()
	model.UpdateTime = time.Now()

	nullTime := sql.NullTime{ // 插入时间为Null
		Time:  time.Now(),
		Valid: false,
	}

	model.BirthTime = nullTime
	fmt.Println(model)

}
