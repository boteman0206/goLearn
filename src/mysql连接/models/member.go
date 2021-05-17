package models

type Member struct {
	Email    string `xorm:"VARCHAR(260)"`
	Sex      string `xorm:"VARCHAR(260)"`
	UserId   int    `xorm:"not null pk autoincr INT(11)"`
	Username string `xorm:"VARCHAR(260)"`
}
