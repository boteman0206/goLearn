package model

import (
	"time"
)

type StoreBusinessSetup struct {
	Id                int       `xorm:"not null pk autoincr INT(11)"`
	FinanceCode       string    `xorm:"not null comment('财务编码') VARCHAR(50)"`
	ChannelId         int       `xorm:"not null comment('渠道id') INT(10)"`
	BusinessStatus    int       `xorm:"not null comment('营业状态1:营业 2:闭店') INT(1)"`
	Businessdate      string    `xorm:"default 'NULL' comment('营业日期（星期一**星期日分别为1**7）') VARCHAR(50)"`
	Advanceorder      int       `xorm:"default NULL comment('预订单设置（true:休息时间支持预订，false:休息时间不支持预订）') TINYINT(1)"`
	Advancedates      string    `xorm:"default 'NULL' comment('接受预定日期（json数组，例[{0,0},{0,1}]）') VARCHAR(256)"`
	Advanceremindtime int       `xorm:"default NULL comment('预订单提醒：送到时间前xx分钟提醒备货') INT(10)"`
	Notice            string    `xorm:"default '''' comment('店铺公告') VARCHAR(256)"`
	Mobile            string    `xorm:"default '''' comment('店铺电话') VARCHAR(128)"`
	Qrcode            string    `xorm:"default '''' comment('店铺二维码') VARCHAR(128)"`
	BusinessCategory1 string    `xorm:"default '''' comment('经营种类1') VARCHAR(56)"`
	BusinessCategory2 string    `xorm:"default 'NULL' comment('经营种类2') VARCHAR(56)"`
	Image             string    `xorm:"default '''' comment('店铺头像') VARCHAR(128)"`
	Businesssystem    string    `xorm:"default '''' comment('业务系统') VARCHAR(56)"`
	Businesssystemid  string    `xorm:"default '''' comment('业务系统ID') VARCHAR(56)"`
	Invoice           int       `xorm:"not null comment('是否支持开具发票(true:支持 false:不支持)') TINYINT(1)"`
	AutoPrint         int       `xorm:"not null comment('是否自动打印(true:自动打印 false:不自动打印)') TINYINT(1)"`
	ReturnAddress     string    `xorm:"default '''' comment('退货地址') VARCHAR(128)"`
	IsSelfLifting     int       `xorm:"default 0 comment('是否支持自提(true:支持 false:不支持)') TINYINT(1)"`
	IsSelfLiftingApp  int       `xorm:"default 0 comment('小程序是否支持自提，1开启，0关闭') TINYINT(1)"`
	StockUpTime       int       `xorm:"default NULL comment('备货时长 单位分钟 不能小于0') INT(11)"`
	Miniappcode       string    `xorm:"default '''' comment('店铺小程序码') VARCHAR(255)"`
	Updatetime        time.Time `xorm:"default 'current_timestamp()' comment('最后更新时间') DATETIME"`
}


//IsSelfLiftingApp  int       `xorm:"default 0 comment('小程序是否支持自提，1开启，0关闭') TINYINT(1)"`
