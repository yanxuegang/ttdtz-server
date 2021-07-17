package models

import (
	"fmt"
	"juliang/global"
	"juliang/pkg/setting"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (map[string]*gorm.DB, error) {
	DBs := make(map[string]*gorm.DB)
	for _, conn := range databaseSetting.DBConns {
		db, err := gorm.Open(conn.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			conn.Username,
			conn.Password,
			conn.Host,
			conn.DBName,
			conn.Charset,
			conn.ParseTime,
		))
		if err != nil {
			panic(fmt.Sprintf("[DB][Exception] %s database connect failed: %s ...\n", conn.DBName, err))
		}
		log.Printf("[DB][Success] %s database connected ...", conn.DBName)
		if global.GlobalConfig.Server.RunMode == "debug" {
			db.LogMode(true)
		}
		db.SingularTable(true)
		db.DB().SetMaxIdleConns(conn.MaxIdleConns)
		db.DB().SetMaxOpenConns(conn.MaxOpenConns)

		DBs[conn.DBName] = db
	}
	return DBs, nil
}
