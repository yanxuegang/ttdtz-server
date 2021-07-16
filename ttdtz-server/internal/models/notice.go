package models

import "time"

type Notice struct {
	Id         int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	Title      string    `gorm:"column:title" db:"title" json:"title" form:"title"`                         //标题
	Type       int       `gorm:"column:type" db:"type" json:"type" form:"type"`                             //1 开服更新 2 停服更新
	NoticeTest string    `gorm:"column:notice_test" db:"notice_test" json:"notice_test" form:"notice_test"` //公告详情
	Level      int16     `gorm:"column:level" db:"level" json:"level" form:"level"`                         //水平格式 0左1中2右
	Vertical   int16     `gorm:"column:vertical" db:"vertical" json:"vertical" form:"vertical"`             //垂直格式 0上1中2下
	CreatedAt  time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`     //添加时间
	IsShow     int16     `gorm:"column:is_show" db:"is_show" json:"is_show" form:"is_show"`                 //1 显示 2删除
}
