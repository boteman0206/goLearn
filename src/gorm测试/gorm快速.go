package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	//"time"

	//"gorm.io/gorm/logger"

	//"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// TableName get sql table name.获取数据库表名
func (m *Product) TableName() string {
	return "product"
}

func main() {

	dns := "root:@(10.1.1.245:3306)/gorm_test_pw?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建连接池
	//sqlDb, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	//sqlDb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//sqlDb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//sqlDb.SetConnMaxLifetime(time.Hour)

	//db.Logger.LogMode(logger.New())

	var data []string
	db.Raw("show tables").Scan(&data)
	fmt.Println("tables: ", data)

	//db.AutoMigrate(&Product{}) // 创建表  默认TableName()返回的字段名称

	//db.Create(&Product{Code:"D42", Price:12})
	product := Product{}
	err = db.First(&product, 1).Error
	fmt.Println(JsonTostring(product), err)
	err = db.First(&product, "id = ?", 1).Error
	fmt.Println(JsonTostring(product), err)

	// 更新update单个字段
	db.Model(&product).Update("price", 899)

	// 更新多个字段 // 仅更新非零值字段
	db.Model(&product).Updates(Product{Price: 10, Code: "DD1"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)

}

func JsonTostring(i interface{}) string {

	bytes, _ := json.Marshal(i)
	return string(bytes)
}
