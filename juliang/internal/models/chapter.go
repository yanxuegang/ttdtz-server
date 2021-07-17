package models

import "time"

type Chapter struct {
	Id           int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	PlayerId     int       `gorm:"column:player_id" db:"player_id" json:"player_id" form:"player_id"`                     //玩家表id
	ChapterId    int       `gorm:"column:chapter_id" db:"chapter_id" json:"chapter_id" form:"chapter_id"`                 //当前大章节
	Type         int       `gorm:"column:type" db:"type" json:"type" form:"type"`                                         //关卡类型
	ChapterSubId int       `gorm:"column:chapter_sub_id" db:"chapter_sub_id" json:"chapter_sub_id" form:"chapter_sub_id"` //打到最高关卡
	MaxSubId     int       `gorm:"column:max_sub_id" db:"max_sub_id" json:"max_sub_id" form:"max_sub_id"`                 //章节最高关卡
	Number       int       `gorm:"column:number" db:"number" json:"number" form:"number"`                                 //次数(星数)
	UpdatedAt    time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`                 //更新时间
}
