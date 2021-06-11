package model

import (
	"time"
)

type UserInfo struct {
	Id          int       `xorm:"not null pk autoincr comment('主键id') unique INT(11)"`
	Userid      string    `xorm:"default 'NULL' comment('成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节') VARCHAR(64)"`
	Mobile      string    `xorm:"default 'NULL' comment('用户手机号码') VARCHAR(50)"`
	Telephone   string    `xorm:"default 'NULL' comment('座机') VARCHAR(30)"`
	Avatar      string    `xorm:"default '''' comment('头像') TEXT"`
	ThumbAvatar string    `xorm:"default 'NULL' comment('头像缩略图') TEXT"`
	Gender      int       `xorm:"not null default 0 comment('性别，0未知，1男，2女') TINYINT(1)"`
	OpenUserId  string    `xorm:"default 'NULL' comment('全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取') VARCHAR(64)"`
	StaffNo     string    `xorm:"not null comment('员工号') unique VARCHAR(30)"`
	Name        string    `xorm:"default 'NULL' comment('用户姓名') VARCHAR(50)"`
	JobName     string    `xorm:"default 'NULL' comment('职位名称') VARCHAR(30)"`
	JobCode     string    `xorm:"default 'NULL' comment('职位编号') VARCHAR(30)"`
	DutyStatus  int       `xorm:"not null default 1 comment('在职状态，1在职，-1离职') TINYINT(1)"`
	CreateTime  time.Time `xorm:"not null default 'current_timestamp()' comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
}
