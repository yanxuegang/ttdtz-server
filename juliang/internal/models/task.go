package models

import "time"

type Task struct {
	Id        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	PlayerId  int       `gorm:"column:player_id" db:"player_id" json:"player_id" form:"player_id"`     //玩家表id
	Type      int       `gorm:"column:type" db:"type" json:"type" form:"type"`                         //任务类型
	Number    int       `gorm:"column:number" db:"number" json:"number" form:"number"`                 //次数
	Status    int       `gorm:"column:status" db:"status" json:"status" form:"status"`                 //领取状态(0:未领,1:已领)
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //添加时间
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"` //更新时间(领取)
}
