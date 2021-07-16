package models

import "time"

type SignDay struct {
	Id        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	PlayerId  int       `gorm:"column:player_id" db:"player_id" json:"player_id" form:"player_id"`     //用户id
	Days      int       `gorm:"column:days" db:"days" json:"days" form:"days"`                         //签到天数
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //签到日期
}
