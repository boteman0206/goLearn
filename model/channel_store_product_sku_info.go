package model

import (
	"time"
)

type ChannelStoreProductSkuInfo struct {
	Id                 int       `xorm:"not null pk autoincr comment('商品id') INT(11)"`
	ProductId          int       `xorm:"default NULL comment('商品SPU') index INT(11)"`
	SkuId              int       `xorm:"default NULL comment('商品SKUID') unique(uk_sku_info) INT(11)"`
	ChannelId          int       `xorm:"default NULL comment('渠道') unique(uk_sku_info) INT(11)"`
	FinanceCode        string    `xorm:"default 'NULL' comment('财务编码') unique(uk_sku_info) VARCHAR(255)"`
	RetailPrice        int       `xorm:"default NULL comment('建议价格') INT(11)"`
	MarketPrice        int       `xorm:"default NULL comment('市场价') INT(11)"`
	PreposePrice       int       `xorm:"default NULL comment('前置仓价格') INT(11)"`
	StorePrice         int       `xorm:"default NULL comment('门店仓价格') INT(11)"`
	BarCode            string    `xorm:"default 'NULL' comment('条码') VARCHAR(36)"`
	WeightForUnit      float64   `xorm:"default NULL comment('重量') DOUBLE(8,2)"`
	WeightUnit         string    `xorm:"default 'NULL' comment('重量单位') VARCHAR(255)"`
	ThirdSkuIdA8       string    `xorm:"default 'NULL' comment('A8货号') VARCHAR(36)"`
	ThirdSpuIdA8       string    `xorm:"default 'NULL' comment('A8spuid') VARCHAR(36)"`
	ThirdSkuIdZl       string    `xorm:"default 'NULL' comment('子龙货号') VARCHAR(36)"`
	ThirdSpuIdZl       string    `xorm:"default 'NULL' comment('子龙spuid') VARCHAR(36)"`
	SkuvSpecId         int       `xorm:"default NULL comment('规格ID') INT(11)"`
	SkuvSpecValueId    int       `xorm:"default NULL comment('规格值ID') INT(11)"`
	SkuvPic            string    `xorm:"default 'NULL' comment('规格属性图片') VARCHAR(500)"`
	SkuvSpecName       string    `xorm:"default 'NULL' comment('规格名称') VARCHAR(255)"`
	SkuvSpecValueValue string    `xorm:"default 'NULL' comment('规格值') VARCHAR(255)"`
	GroupProductId     int       `xorm:"default NULL comment('组合商品明细商品ID') INT(11)"`
	GroupSkuId         int       `xorm:"default 0 comment('组合商品明细商品SKUID') unique(uk_sku_info) INT(11)"`
	GroupCount         int       `xorm:"default NULL comment('组合商品数量') INT(11)"`
	GroupProductName   string    `xorm:"default 'NULL' comment('组合商品名称') VARCHAR(255)"`
	GroupDiscountType  int       `xorm:"default NULL comment('折扣类型（1按折扣优惠，2按固定价格优惠）') INT(11)"`
	GroupDiscountValue int       `xorm:"default NULL comment('折扣值（当折扣类型为1时，存百分比。折扣类型为2时，存具体设置的价格。）(单位分)') INT(11)"`
	GroupMarketPrice   int       `xorm:"default NULL comment('组合商品明细市场价') INT(11)"`
	GroupProductType   int       `xorm:"default NULL comment('组合商品明细类型') INT(11)"`
	SnapshotMainId     int       `xorm:"default NULL comment('快照主表ID') INT(11)"`
	CreateDate         time.Time `xorm:"default 'current_timestamp()' comment('创建时间') DATETIME"`
	UpdateDate         time.Time `xorm:"default 'current_timestamp()' comment('修改时间') DATETIME"`
	LastOperator       string    `xorm:"default 'NULL' comment('最后修改人') VARCHAR(100)"`
	LastFunction       string    `xorm:"default 'NULL' comment('最后修改函数') VARCHAR(100)"`
}
