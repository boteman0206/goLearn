package main

import (
	"database/sql/driver"
	"fmt"
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

}
