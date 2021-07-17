package models

type OpenidMap struct {
	Openid   string `gorm:"primary_key"`
	System   uint8  `gorm:"primary_key"`
	PlayerId uint64
}

func (m OpenidMap) TableName() string {
	return "openid_map"
}
