package models

import "time"

type TaskList struct {
	Id        int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	PlayerId  int       `gorm:"column:player_id" db:"player_id" json:"player_id" form:"player_id"`     //用户id
	Type      int16     `gorm:"column:type" db:"type" json:"type" form:"type"`                         //1:领取签到红包 2:闯关三次 3:下载app三次 4:评分三星以上5次 5:观看视频5次
	Num       int16     `gorm:"column:num" db:"num" json:"num" form:"num"`                             //任务次数
	Money     int       `gorm:"column:money" db:"money" json:"money" form:"money"`                     //红包币
	Name      string    `gorm:"column:name" db:"name" json:"name" form:"name"`                         //任务名称
	Status    int16     `gorm:"column:status" db:"status" json:"status" form:"status"`                 //任务状态1未完成 2 已完成
	GetState  int16     `gorm:"column:get_state" db:"get_state" json:"get_state" form:"get_state"`     //领取奖励状态 1:未领取 2:已领取
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //任务时间
}
