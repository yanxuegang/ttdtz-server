package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
}

func NewDBEngine() (*gorm.DB, error) {
	DB := &gorm.DB{}
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		"root",
		"c6067720ecd29244",
		"127.0.0.1:3306",
		"app_line",
		"utf8",
		true,
	))
	if err != nil {
		panic(fmt.Sprintf("[DB][Exception] %s database connect failed: %s ...\n", "app_line", err))
	}
	log.Printf("[DB][Success] %s database connected ...", "app_line")

	db.LogMode(true)

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(30)

	DB = db

	return DB, nil
}
