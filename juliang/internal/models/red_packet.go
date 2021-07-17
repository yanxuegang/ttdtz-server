package models

import "time"

type RedPacket struct {
	Id        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	PlayerId  int       `gorm:"column:player_id" db:"player_id" json:"player_id" form:"player_id"`     //玩家id
	Money     int       `gorm:"column:money" db:"money" json:"money" form:"money"`                     //随机红包币
	Token     string    `gorm:"column:token" db:"token" json:"token" form:"token"`                     //签名值
	Time      string    `gorm:"column:time" db:"time" json:"time" form:"time"`                         //签名时间
	LogDate   time.Time `gorm:"column:log_date" db:"log_date" json:"log_date" form:"log_date"`         //日期
	Status    int       `gorm:"column:status" db:"status" json:"status" form:"status"`                 //0:未领取,1:领取
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //记录生成时间
}
