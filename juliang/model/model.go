package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB = &gorm.DB{}
)

//监测数据sqlmodel
type Translations struct {
	Id          int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Aid         string    `gorm:"column:aid" db:"aid" json:"aid" form:"aid"`
	ConvertId   string    `gorm:"column:convert_id" db:"convert_id" json:"convert_id" form:"convert_id"`
	RequestId   string    `gorm:"column:request_id" db:"request_id" json:"request_id" form:"request_id"`
	Imei        string    `gorm:"column:imei" db:"imei" json:"imei" form:"imei"`
	Idfa        string    `gorm:"column:idfa" db:"idfa" json:"idfa" form:"idfa"`
	Androidid   string    `gorm:"column:androidid" db:"androidid" json:"androidid" form:"androidid"`
	Oaid        string    `gorm:"column:oaid" db:"oaid" json:"oaid" form:"oaid"`
	OaidMd5     string    `gorm:"column:oaid_md5" db:"oaid_md5" json:"oaid_md5" form:"oaid_md5"`
	Os          int       `gorm:"column:os" db:"os" json:"os" form:"os"`
	Mac         string    `gorm:"column:mac" db:"mac" json:"mac" form:"mac"`
	Mac1        string    `gorm:"column:mac1" db:"mac1" json:"mac1" form:"mac1"`
	Ip          string    `gorm:"column:ip" db:"ip" json:"ip" form:"ip"`
	Ua          string    `gorm:"column:ua" db:"ua" json:"ua" form:"ua"`
	Geo         string    `gorm:"column:geo" db:"geo" json:"geo" form:"geo"`
	Ts          time.Time `gorm:"column:ts" db:"ts" json:"ts" form:"ts"`
	CallbackUrl string    `gorm:"column:callback_url" db:"callback_url" json:"callback_url" form:"callback_url"`
	Callback    string    `gorm:"column:callback" db:"callback" json:"callback" form:"callback"`
	Model       string    `gorm:"column:model" db:"model" json:"model" form:"model"`
	Status      int       `gorm:"column:status" db:"status" json:"status" form:"status"`
}

func (m *Translations) Create() error {
	//todo redis save
	return DB.Create(m).Error
}

func (m *Translations) Delete() error {
	//todo redis save
	return DB.Delete(m).Error
}

func (m *Translations) Update(attrs ...interface{}) error {
	//todo redis save
	rowsAffected := DB.Model(m).Update(attrs...).RowsAffected
	if rowsAffected == 0 {
		log.Println("[PLAYER][WARNING] No Content To Update.")
		return nil
	}
	return nil
}

type Orders struct {
	Id                 int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Appid              string    `gorm:"column:appid" db:"appid" json:"appid" form:"appid"`
	MchId              string    `gorm:"column:mch_id" db:"mch_id" json:"mch_id" form:"mch_id"`
	DeviceInfo         string    `gorm:"column:device_info" db:"device_info" json:"device_info" form:"device_info"`
	NonceStr           string    `gorm:"column:nonce_str" db:"nonce_str" json:"nonce_str" form:"nonce_str"`
	Openid             string    `gorm:"column:openid" db:"openid" json:"openid" form:"openid"`
	TotalFee           float64   `gorm:"column:total_fee" db:"total_fee" json:"total_fee" form:"total_fee"`
	SettlementTotalFee int       `gorm:"column:settlement_total_fee" db:"settlement_total_fee" json:"settlement_total_fee" form:"settlement_total_fee"`
	FeeType            string    `gorm:"column:fee_type" db:"fee_type" json:"fee_type" form:"fee_type"`
	CashFee            int       `gorm:"column:cash_fee" db:"cash_fee" json:"cash_fee" form:"cash_fee"`
	TransactionId      string    `gorm:"column:transaction_id" db:"transaction_id" json:"transaction_id" form:"transaction_id"`
	OutTradeNo         string    `gorm:"column:out_trade_no" db:"out_trade_no" json:"out_trade_no" form:"out_trade_no"`
	TimeEnd            time.Time `gorm:"column:time_end" db:"time_end" json:"time_end" form:"time_end"`
	Status             int       `gorm:"column:status" db:"status" json:"status" form:"status"`
	CreatedAt          time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //提现时间
}

func (m *Orders) Create() error {
	//todo redis save
	return DB.Create(m).Error
}

func (m *Orders) Delete() error {
	//todo redis save
	return DB.Delete(m).Error
}

func (m *Orders) Update(attrs ...interface{}) error {
	//todo redis save
	rowsAffected := DB.Model(m).Update(attrs...).RowsAffected
	if rowsAffected == 0 {
		log.Println("[PLAYER][WARNING] No Content To Update.")
		return nil
	}
	return nil
}

type Model struct {
}

func NewDBEngine() (*gorm.DB, error) {
	DB = &gorm.DB{}
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		"root",
		"root", //"e90e3a8a2ec08eeb",
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
