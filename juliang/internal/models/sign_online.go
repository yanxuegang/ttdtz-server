package models

import "time"

type SignOnline struct {
	Id        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	PlayerId  int       `gorm:"column:player_id" db:"player_id" json:"player_id" form:"player_id"`     //玩家id
	Type      int       `gorm:"column:type" db:"type" json:"type" form:"type"`                         //在线类型(分钟)
	Second    int       `gorm:"column:second" db:"second" json:"second" form:"second"`                 //在线时间(签到时总在线时长)
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //记录时间
}
