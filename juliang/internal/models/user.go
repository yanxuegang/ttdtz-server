package models

import "time"

type User struct {
	Id              int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	OpenId          string    `gorm:"column:open_id" db:"open_id" json:"open_id" form:"open_id"`                                         //微信openId
	Password        string    `gorm:"column:password" db:"password" json:"password" form:"password"`                                     //登录密码
	Type            string    `gorm:"column:type" db:"type" json:"type" form:"type"`                                                     //来源
	UnionId         string    `gorm:"column:union_id" db:"union_id" json:"union_id" form:"union_id"`                                     //微信union_id
	ServerId        int       `gorm:"column:server_id" db:"server_id" json:"server_id" form:"server_id"`                                 //区服id
	CreatedAt       time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                             //注册时间
	LoginAt         time.Time `gorm:"column:login_at" db:"login_at" json:"login_at" form:"login_at"`                                     //登录时间
	NoticeVersionAt time.Time `gorm:"column:notice_version_at" db:"notice_version_at" json:"notice_version_at" form:"notice_version_at"` //公告版本弹一次
	NoticeDailyAt   time.Time `gorm:"column:notice_daily_at" db:"notice_daily_at" json:"notice_daily_at" form:"notice_daily_at"`         //每日首次必弹
	IsWhiteType     int       `gorm:"column:is_white_type" db:"is_white_type" json:"is_white_type" form:"is_white_type"`                 //是否白名单类型(0:不是,1:是)
	Ip              string    `gorm:"column:ip" db:"ip" json:"ip" form:"ip"`                                                             //ip地址
}
