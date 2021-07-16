package global

import (
	"github.com/jinzhu/gorm"
)

var (
	DBEngine map[string]*gorm.DB
)

func GetDB(name string) *gorm.DB {
	if db, ok := DBEngine[name]; ok {
		return db
	}
	return nil
}
