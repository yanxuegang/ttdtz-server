package models

import "time"

type Payment struct {
	Id             int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	PlayerId       int       `gorm:"column:player_id" db:"player_id" json:"player_id" form:"player_id"`                             //玩家id
	Amount         float64   `gorm:"column:amount" db:"amount" json:"amount" form:"amount"`                                         //提现金额(分)
	CreatedAt      time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                         //提现时间
	PartnerTradeNo string    `gorm:"column:partner_trade_no" db:"partner_trade_no" json:"partner_trade_no" form:"partner_trade_no"` //商户订单号
	PaymentNo      string    `gorm:"column:payment_no" db:"payment_no" json:"payment_no" form:"payment_no"`                         //微信付款单号
	PaymentTime    string    `gorm:"column:payment_time" db:"payment_time" json:"payment_time" form:"payment_time"`                 //企业付款成功时间
}
