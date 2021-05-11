package main

type UserInfo struct {
	userId   string `xorm:"not null pk VARCHAR(50)"`
	userName string `xorm:"default 'NULL' comment('批次号') VARCHAR(50)"`
	sex      string `xorm:"default 'NULL' comment(性别，男女)"`
	email    string `xorm:"default 'NULL' comment(邮箱账号)"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

type Member struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Person struct {
	Id   int    `db:"id" xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
