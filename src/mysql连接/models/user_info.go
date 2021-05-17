package models

type UserInfo struct {
	Email    string `xorm:"default 'NULL' comment('邮箱账号') VARCHAR(255)"`
	Sex      string `xorm:"default 'NULL' comment('性别，男女') VARCHAR(255)"`
	UserId   string `xorm:"not null pk VARCHAR(50)"`
	UserName string `xorm:"default 'NULL' comment('批次号') VARCHAR(50)"`
}
