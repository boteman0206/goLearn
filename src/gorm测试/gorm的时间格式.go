package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"strings"
	"time"
)

// 库存变更日志
type StockLog struct {
	Id          int64    `json:"id"`           // 主键id,自增
	BillTime    JSONTime `json:"bill_time"`    // 单据日期
	BillSn      string   `json:"bill_sn"`      // 单据编
	ProductName string   `json:"product_name"` // 商品名称

	AlterNum     int32    `json:"alter_num"`            // 变动数量
	LatestNum    int32    `json:"latest_num"`           // 新库存数量(原库存数量+变动数量)
	LatestAmount int64    `json:"latest_amount"`        // 新库存成本(原库存成本+变动库存成本)
	LatestPrice  int64    `json:"latest_price"`         // 新成本单价(新库存成本/新库存数量)
	OldNum       int32    `json:"old_num"`              // 原库存数量
	OldPrice     int64    `json:"old_price"`            // 原成本单价
	OldAmount    int64    `json:"old_amount"`           // 原库存成本
	AlterPrice   int64    `json:"alter_price"`          // 变动库存单价
	AlterAmount  int64    `json:"alter_amount"`         // 变动库存成本
	AlterTime    JSONTime `gorm:"->" json:"alter_time"` // 创建时间
}

func (m StockLog) TableName() string {
	return "stock_log"
}

const TimeFormat = "2006-01-02 15:04:05"

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = JSONTime{Time: time.Time{}}
		return
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), loc)
	*t = JSONTime{Time: now}
	return
}

// MarshalJSON on JSONTime format Time field with Y-m-d H:i:s
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func main() {

	location, _ := time.ParseInLocation(TimeFormat, "2022-01-12 11:23:12", time.Local)

	fmt.Println(time.Now().String())

	jsonTime := JSONTime{location}

	fmt.Println(jsonTime)

	//myTime := MyTime(cast.ToTime(""))

}

//BaseModel 基础结构体 信息信息
type BaseModel struct {
	CreateTime MyTime `gorm:"comment:'创建时间';type:timestamp;";json:"createTime"`
	UpdateTime MyTime `gorm:"comment:'修改时间';type:timestamp;";json:"updateTime"`
	Remark     string `gorm:"comment:'备注'";json:"remark"`
}

//MyTime 自定义时间
type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
