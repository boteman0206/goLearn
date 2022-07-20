package model

import (
	"time"
)

type ChannelStoreProductSnapshotMain struct {
	Id                   int       `xorm:"not null pk autoincr comment('商品id') INT(11)"`
	ChannelId            int       `xorm:"not null pk default 0 comment('渠道id(1-阿闻，2-美团，3-饿了么，4-京东到家)') unique(idx_only) INT(11)"`
	ProductId            int       `xorm:"default NULL comment('商品SPU') unique(idx_only) INT(11)"`
	FinanceCode          string    `xorm:"default 'NULL' comment('财务编码') unique(idx_only) VARCHAR(50)"`
	CategoryId           int       `xorm:"default NULL comment('分类id') INT(11)"`
	BrandId              int       `xorm:"default NULL comment('品牌id') INT(11)"`
	Name                 string    `xorm:"default 'NULL' comment('商品名称') VARCHAR(255)"`
	Code                 string    `xorm:"default 'NULL' comment('商品编号') VARCHAR(36)"`
	BarCode              string    `xorm:"default 'NULL' comment('商品条码') VARCHAR(36)"`
	Pic                  string    `xorm:"default 'NULL' comment('商品图片（多图）') VARCHAR(1000)"`
	SellingPoint         string    `xorm:"default 'NULL' comment('商品卖点') VARCHAR(200)"`
	Video                string    `xorm:"default 'NULL' comment('商品视频地址') VARCHAR(500)"`
	ContentPc            string    `xorm:"default 'NULL' comment('电脑端详情内容') TEXT"`
	ContentMobile        string    `xorm:"default 'NULL' comment('手机端详情内容') TEXT"`
	ProductType          int       `xorm:"default NULL comment('商品类别（1-实物商品，2-虚拟商品，3-组合商品') INT(11)"`
	IsUse                int       `xorm:"default 0 comment('商品是否被使用过（认领或者其它第三方使用，否则不能被删除商品本身及SKU）') INT(11)"`
	ChannelCategoryId    int       `xorm:"not null default 0 comment('渠道的分类id') INT(11)"`
	ChannelName          string    `xorm:"default 'NULL' comment('渠道名称（美团，饿了么，阿闻，京东）') VARCHAR(50)"`
	LastEditUser         string    `xorm:"default 'NULL' comment('最后编辑用户') VARCHAR(50)"`
	ChannelTagId         string    `xorm:"default ''0'' comment('渠道商品类目id') VARCHAR(50)"`
	CategoryName         string    `xorm:"default 'NULL' comment('分类名称') VARCHAR(100)"`
	ChannelCategoryName  string    `xorm:"default 'NULL' comment('渠道分类名称') VARCHAR(100)"`
	IsRecommend          int       `xorm:"not null default 0 comment('是否是推荐商品,1是0否') TINYINT(1)"`
	BrandName            string    `xorm:"default 'NULL' comment('品牌名称') VARCHAR(255)"`
	GroupType            int       `xorm:"not null default 0 comment('组合类型(1:实实组合,2:虚虚组合,3.虚实组合)') INT(11)"`
	TermType             int       `xorm:"not null default 0 comment('只有虚拟商品才有值(1.有效期至多少  2.有效期天数)') INT(11)"`
	TermValue            int       `xorm:"not null default 0 comment('如果term_type=1 存：时间戳  如果term_type=2 存多少天') INT(11)"`
	UseRange             string    `xorm:"not null default '''' comment('商品应用范围（1电商，2前置仓，3门店仓）,字符串拼接') VARCHAR(100)"`
	VirtualInvalidRefund int       `xorm:"not null default 0 comment('是否支持过期退款 1：是  0：否') INT(11)"`
	CreateDate           time.Time `xorm:"default 'current_timestamp()' comment('创建时间') DATETIME"`
	UpdateDate           time.Time `xorm:"default 'current_timestamp()' comment('修改时间') DATETIME"`
	LastOperator         string    `xorm:"default 'NULL' comment('最后修改人') VARCHAR(100)"`
	LastFunction         string    `xorm:"default 'NULL' comment('最后修改函数') VARCHAR(100)"`
}
