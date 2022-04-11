package main

import "time"

type User struct {
	Name string `gorm:"<-:create"`          // 允许读和创建
	Name string `gorm:"<-:update"`          // 允许读和更新
	Name string `gorm:"<-"`                 // 允许读和写（创建和更新）
	Name string `gorm:"<-:false"`           // 允许读，禁止写
	Name string `gorm:"->"`                 // 只读（除非有自定义配置，否则禁止写）
	Name string `gorm:"->;<-:create"`       // 允许读和写
	Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
	Name string `gorm:"-"`                  // 读写操作均会忽略该字段
}

type User struct {
	CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
	UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	Updated   int64     `gorm:"autoUpdateTime:nano"`  // 使用时间戳填纳秒数充更新时间
	Updated   int64     `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created   int64     `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
}

//字段嵌入

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

// 等效于
type Blog struct {
	ID      int64
	Name    string
	Email   string
	Upvotes int32
}

/**
todo 想要正确的处理 time.Time ，您需要带上 parseTime 参数， (更多参数) 要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4 查看 此文章 获取详情


*/
