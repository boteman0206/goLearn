package models

type Place struct {
	City    string `xorm:"VARCHAR(200)"`
	Country string `xorm:"VARCHAR(200)"`
	Telcode int    `xorm:"INT(11)"`
}
