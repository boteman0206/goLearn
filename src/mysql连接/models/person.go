package models

type Person struct {
	Age  int    `xorm:"INT(11)"`
	Id   int    `xorm:"not null pk autoincr INT(11)"`
	Name string `xorm:"VARCHAR(20)"`
	Sex  string `xorm:"VARCHAR(12)"`
}
