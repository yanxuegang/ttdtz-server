package models

import (
	"log"
	"time"
	"ttdtz-server/global"

	"github.com/jinzhu/gorm"
)

// 用户表
type AclUsers struct {
	Id        uint      `gorm:"column:id" db:"id" json:"id" form:"id"`
	Platform  string    `gorm:"column:platform" db:"platform" json:"platform" form:"platform"`         //平台
	Username  string    `gorm:"column:username" db:"username" json:"username" form:"username"`         //用户名
	Password  string    `gorm:"column:password" db:"password" json:"password" form:"password"`         //密码
	Email     string    `gorm:"column:email" db:"email" json:"email" form:"email"`                     //邮箱
	IsAdmin   int       `gorm:"column:is_admin" db:"is_admin" json:"is_admin" form:"is_admin"`         //是否管理员(0:不是,1:是)
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //添加时间
}

func (m *AclUsers) GetDB() *gorm.DB {
	db := global.GetDB("app_line")
	if db == nil {
		log.Println("[Error] Cannot connect DB")
		return nil
	}
	return db
}

func (m *AclUsers) Create() error {
	//todo redis save
	return m.GetDB().Create(m).Error
}

func (m *AclUsers) Update(attrs ...interface{}) error {
	//todo redis save
	rowsAffected := m.GetDB().Model(m).Update(attrs...).RowsAffected
	if rowsAffected == 0 {
		log.Println("[PLAYER][WARNING] No Content To Update.")
		return nil
	}
	return nil
}
