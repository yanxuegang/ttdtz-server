package models

import "time"

type SourceJuliang struct {
	Id          int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	OpenId      string    `gorm:"column:open_id" db:"open_id" json:"open_id" form:"open_id"`                     //微信openId
	Oaid        string    `gorm:"column:oaid" db:"oaid" json:"oaid" form:"oaid"`                                 //Android Q版本的oaid原值
	ImeiMd5     string    `gorm:"column:imei_md5" db:"imei_md5" json:"imei_md5" form:"imei_md5"`                 //安卓系统imei的md5摘要
	Aid         string    `gorm:"column:aid" db:"aid" json:"aid" form:"aid"`                                     //广告计划id
	Os          string    `gorm:"column:os" db:"os" json:"os" form:"os"`                                         //操作系统
	CallbackUrl string    `gorm:"column:callback_url" db:"callback_url" json:"callback_url" form:"callback_url"` //回调地址
	CreatedAt   time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
}
