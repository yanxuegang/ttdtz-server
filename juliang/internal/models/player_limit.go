package models

import "time"

type PlayerLimit struct {
	Id        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	OpenId    string    `gorm:"column:open_id" db:"open_id" json:"open_id" form:"open_id"`             //禁言账号
	CreatedId time.Time `gorm:"column:created_id" db:"created_id" json:"created_id" form:"created_id"` //添加时间
}
