package models

type Service struct {
	Id      int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name    string `gorm:"column:name" db:"name" json:"name" form:"name"`             //服务器名
	Path    string `gorm:"column:path" db:"path" json:"path" form:"path"`             //游戏地址(服务器端请求地址)
	Type    int    `gorm:"column:type" db:"type" json:"type" form:"type"`             //区服类型(0:普通 1:新开 2:推荐 3:新开推荐)
	Status  int    `gorm:"column:status" db:"status" json:"status" form:"status"`     //服务器状态(0:关服 1:顺畅 2:拥挤 3:爆满,4:维护)
	Channel string `gorm:"column:channel" db:"channel" json:"channel" form:"channel"` //渠道来源
	Zoneno  int    `gorm:"column:zoneno" db:"zoneno" json:"zoneno" form:"zoneno"`     //区服编号(x区)
	Times   int    `gorm:"column:times" db:"times" json:"times" form:"times"`         //开服时间戳(秒)
	Normal  int    `gorm:"column:normal" db:"normal" json:"normal" form:"normal"`     //正常人数(区服玩家数)
	Current int    `gorm:"column:current" db:"current" json:"current" form:"current"` //当前人数
	Signkey string `gorm:"column:signkey" db:"signkey" json:"signkey" form:"signkey"` //签名key值
	Version string `gorm:"column:version" db:"version" json:"version" form:"version"` //版本号
}
