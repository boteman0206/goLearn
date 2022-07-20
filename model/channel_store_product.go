package model

import (
	"time"
)

type ChannelStoreProduct struct {
	Id                  int       `xorm:"not null pk autoincr INT(11)"`
	ChannelId           int       `xorm:"default NULL comment('渠道id') index index(idx_channelid_updatedate) index(idx_chn_finanche_product) INT(11)"`
	FinanceCode         string    `xorm:"default 'NULL' comment('财务编码') index(idx_chn_finanche_product) index VARCHAR(50)"`
	ProductId           int       `xorm:"default NULL comment('商品id') index(idx_chn_finanche_product) index INT(11)"`
	ChannelCategoryId   int       `xorm:"not null default 0 comment('分类id') INT(11)"`
	IsRecommend         int       `xorm:"not null default 0 comment('是否是推荐商品,1是0否') TINYINT(1)"`
	UpDownState         int       `xorm:"default NULL comment('上下架状态（1-上架，0-下架）') INT(11)"`
	CreateDate          time.Time `xorm:"default 'current_timestamp()' comment('创建时间') DATETIME"`
	UpdateDate          time.Time `xorm:"default 'current_timestamp()' comment('更新时间') index(idx_channelid_updatedate) index DATETIME"`
	SnapshotId          int       `xorm:"default NULL comment('快照id，最后一次上架更新用的是哪个快照') index INT(11)"`
	ChannelCategoryName string    `xorm:"default 'NULL' comment('渠道分类名称') VARCHAR(100)"`
	Name                string    `xorm:"default 'NULL' comment('商品名称') VARCHAR(255)"`
	SkuId               int       `xorm:"default NULL comment('商品sku') INT(11)"`
	MarketPrice         int       `xorm:"default NULL comment('市场价') INT(11)"`
	SalesVolume         int       `xorm:"default 0 comment('销量') INT(11)"`
	HasStock            int       `xorm:"default 0 comment('有无库存') INT(11)"`
	DownType            int       `xorm:"default 0 comment('下架类型：0：之前的默认数据 1： 目前只有7天无库存下架') INT(11)"`
}
