package model

import (
	"time"
)

type ProductThirdInfo struct {
	Id         int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Vid        string    `xorm:"default 'NULL' comment('微盟的vid') VARCHAR(50)"`
	SpuId      string    `xorm:"default 'NULL' comment('第三方的spu_id') VARCHAR(50)"`
	SkuId      string    `xorm:"default 'NULL' comment('第三方的sku_id') VARCHAR(50)"`
	ThirdSkuId string    `xorm:"default 'NULL' comment('oms系统的货号对应的第三方货号') VARCHAR(50)"`
	Source     int       `xorm:"default NULL comment('第三方来源目前只有微盟(1：微盟)') INT(11)"`
	UpdateDate time.Time `xorm:"default 'current_timestamp()' DATETIME"`
}
