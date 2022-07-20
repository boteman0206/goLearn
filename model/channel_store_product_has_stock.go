package model

import (
	"time"
)

type ChannelStoreProductHasStock struct {
	Id                    int       `xorm:"not null pk autoincr unique INT(11)"`
	ChannelStoreProductId int       `xorm:"not null comment('channel_store_product的关联id') index INT(11)"`
	WarehouseId           int       `xorm:"default NULL comment('仓库表的id') index INT(11)"`
	ChannelId             int       `xorm:"default NULL comment('渠道id') INT(11)"`
	SkuId                 int       `xorm:"default NULL comment('sku_id') index INT(11)"`
	FinanceCode           string    `xorm:"default 'NULL' comment('财务编码') VARCHAR(50)"`
	ProductId             int       `xorm:"default NULL comment('商品id') INT(11)"`
	CreateTime            time.Time `xorm:"default 'current_timestamp()' comment('创建时间') DATETIME"`
	UpdateTime            time.Time `xorm:"default 'current_timestamp()' comment('更新时间') DATETIME"`
	HasStock              int       `xorm:"default 0 comment('有无库存') INT(11)"`
	HasStockUp            int       `xorm:"default 0 comment('有库存，并上架 1。') INT(11)"`
	UpDownState           int       `xorm:"default 0 comment('上下架状态 1 上架 0 下架') INT(11)"`
}
