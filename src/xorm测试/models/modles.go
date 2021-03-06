package models

type UserInfo struct {
	Id       int    `xorm:"pk"`
	UserId   string `xorm:"not null pk VARCHAR(50)"`
	UserName string `xorm:"default 'NULL' comment('批次号') VARCHAR(50)"`
	Sex      string `xorm:"default 'NULL' comment(性别，男女)"`
	Email    string `xorm:"default 'NULL' comment(邮箱账号)"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

type Member struct {
	UserId   int    `db:"user_id" xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

// todo 如果表里面有相同的字段则无法展示，在join的时候，可以在字段前面加上字段去区分
type Person struct {
	Id   int    `db:"id" xorm:"pk"`
	Name string `db:"name"`
	Age  int    `db:"age"`

	Sex string `db:"sex"`
}
