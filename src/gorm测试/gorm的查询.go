package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm/logger"
	"time"

	//"gorm.io/gorm/logger"

	//"encoding/json"
	//"fmt"
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
	config := &gorm.Config{}
	config.Logger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dns), config)
	if err != nil {
		panic("failed to connect database")
	}

	var product Product

	err = db.First(&product, 10).Error

	is := errors.Is(err, gorm.ErrRecordNotFound) // todo 这里会返回 record not found的错误

	fmt.Print(product, is)

	var data []Product
	db.Select("*").Find(&data)

	fmt.Println("debug：", data)

	now := time.Now()
	fmt.Println(now)

	location, err := time.ParseInLocation("2006-01-02", "2022-04-07 12:31:39", time.Local)
	//year, month, day := location.Date()

	fmt.Println(err, location)

}
