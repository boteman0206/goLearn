package model

type ChannelStoreProductPrices struct {
	FinanceCode     string `xorm:"default 'NULL' index VARCHAR(50)"`
	ProductId       int    `xorm:"default NULL index INT(11)"`
	SkuId           int    `xorm:"default NULL index INT(11)"`
	Price           string `xorm:"default 'NULL' index VARCHAR(50)"`
	Warehousecatory int    `xorm:"default NULL index INT(11)"`
}
